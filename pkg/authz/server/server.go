package server

import (
	"context"
	b64 "encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"regexp"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/ovotech/kiss/pkg/keyfunc"
	pb "github.com/ovotech/kiss/proto"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

var (
	refreshUnknownKID = true
)

type serverAuthzInterceptor struct {
	jwks            *keyfunc.JWKs
	namespacesKey   string
	namespacesRegex string
	identifierKey   string
	adminNamespace  string
}

type RequestWithMetadata interface {
	GetMetadata() *pb.ClientMeta
}

// struct to hold claims extracted from the token.
type claims struct {
	namespaces []string
	identifier string
}

// Returns a new ServerAuthzInterceptor for validating authorization tokens in client requests.
// jwksURL is used to fetch the JWKS for validating incoming keys. This will refresh if a kid is
// unknown.
// namespacesKey is the key to a list of namespace claims in the token.
// namespacesRegex is used for extracting namespace from the claims. For example, given
// `kaluza:default`, we can usethe regex `kaluza:([1-9a-z-]{1,63})`` to extract the namespace
// `default` from the claim.
// identiferKey is the key to a unique identifier in the claim, for example the email. This is for
// auditing purposes.
func NewServerAuthzInterceptor(
	jwksURL, namespacesKey, namespacesRegex, identifierKey, adminNamespace string,
) *serverAuthzInterceptor {
	jwks, err := keyfunc.Get(jwksURL, keyfunc.Options{RefreshUnknownKID: &refreshUnknownKID})
	if err != nil {
		log.Fatal().Msgf("Failed to get the JWKS from the given URL: %s", err.Error())
	}

	return &serverAuthzInterceptor{
		jwks:            jwks,
		namespacesKey:   namespacesKey,
		namespacesRegex: namespacesRegex,
		identifierKey:   identifierKey,
		adminNamespace:  adminNamespace,
	}
}

// Intercept client requests to validate the token, authorize the request, and keep an audit.
func (i *serverAuthzInterceptor) Unary() grpc.UnaryServerInterceptor {
	return func(
		ctx context.Context,
		req interface{},
		info *grpc.UnaryServerInfo,
		handler grpc.UnaryHandler,
	) (interface{}, error) {
		mdr, ok := req.(RequestWithMetadata)
		if !ok {
			return nil, status.Errorf(codes.InvalidArgument, "missing client metadata")
		}

		// errors if token is invalid (expired, bad signature or unparsable)
		claims, err := i.parseToken(ctx)
		if err != nil {
			return nil, status.Errorf(codes.Unauthenticated, err.Error())
		}

		// errors if user is not authorized to manipulate a secret for the given namespace
		err = i.authorize(claims, mdr.GetMetadata().GetNamespace())
		if err != nil {
			auditLog(
				ctx,
				false,
				info.FullMethod,
				claims.identifier,
				mdr.GetMetadata().Namespace,
			)
			return nil, status.Errorf(codes.PermissionDenied, err.Error())
		}

		auditLog(
			ctx,
			true,
			info.FullMethod,
			claims.identifier,
			mdr.GetMetadata().GetNamespace(),
		)

		newCtx := context.WithValue(ctx, "user", claims.identifier)

		return handler(newCtx, req)
	}
}

// Parses the access token from the context and returns the custom claims, validating the token in
// the process.
// Errors out if token fails validation for whatever reason. If no error, the token's claims can be
// trusted.
func (i *serverAuthzInterceptor) parseToken(ctx context.Context) (*claims, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, errors.New("grpc metadata is not provided")
	}

	authzValues := md["authorization"]
	if len(authzValues) == 0 {
		return nil, errors.New("authorization token is not provided")
	}

	accessToken := authzValues[0]

	// We're only using jwt.Parse() for validation.
	// This is because we want some custom logic for extracting claims
	token, err := jwt.Parse(
		accessToken,
		i.jwks.KeyFunc,
	)
	if err != nil {
		return nil, err
	}

	claims, err := i.getCustomClaims(token)
	if err != nil {
		return nil, err
	}

	return claims, nil
}

// Returns an error if the user's claims is not authorized to perform an action on the
// requestNamespace. This assumes the token has already been validated.
// Note: this function does _not_ validate the token.
func (i *serverAuthzInterceptor) authorize(claims *claims, requestNamespace string) error {
	for _, claimNamespace := range claims.namespaces {
		if claimNamespace == requestNamespace || (i.adminNamespace != "" && i.adminNamespace == claimNamespace) {
			return nil
		}
	}

	return errors.New(
		fmt.Sprintf(
			"user '%s' is not authorized for namespace '%s'",
			claims.identifier,
			requestNamespace,
		),
	)
}

// Extracts claims from a token. We use our own function for this instead of the library's one
// because we want to extract namespaces from strings with some regex, and support configuring the
// keys for the list of namespaces and the user identifier at runtime.
// Note: this function does _not_ validate the token.
func (i *serverAuthzInterceptor) getCustomClaims(token *jwt.Token) (*claims, error) {
	b64Payload := strings.Split(token.Raw, ".")[1]
	strPayload, err := b64.RawStdEncoding.DecodeString(b64Payload)
	if err != nil {
		return nil, err
	}

	var payload map[string]json.RawMessage
	err = json.Unmarshal(strPayload, &payload)
	if err != nil {
		return nil, err
	}

	var identifier string
	if val, ok := payload[i.identifierKey]; ok {
		json.Unmarshal(val, &identifier)
	} else {
		return nil, errors.New(fmt.Sprintf("failed unmarshalling '%s' from token", i.identifierKey))
	}

	var rawNamespaces []string
	if val, ok := payload[i.namespacesKey]; ok {
		json.Unmarshal(val, &rawNamespaces)
	} else {
		return nil, errors.New(fmt.Sprintf("failed unmarshalling '%s' from token", i.namespacesKey))
	}

	// If we don't have a regexp to extract namespaces from the raw namespaces list, we're done
	if i.namespacesRegex == "" {
		return &claims{identifier: identifier, namespaces: rawNamespaces}, nil
	}

	// If we have a regexp set, we're going to run through each namespace in the list and extract
	// the matching group.
	//
	// This allows us to support cases where claims may look like:
	//
	//   "cognito:groups": [
	//     "kaluza:default",
	//     "kaluza:kube-system",
	//   ],
	//
	// and we want to get the namespaces 'default' and 'kube-system'.
	re := regexp.MustCompile(i.namespacesRegex)
	var namespaces []string
	for _, n := range rawNamespaces {
		matches := re.FindStringSubmatch(n)
		if len(matches) != 2 {
			return nil, errors.New(
				fmt.Sprintf(
					"failed to extract namespace from '%s' using regexp `%s`",
					n,
					i.namespacesRegex,
				),
			)
		}
		namespaces = append(namespaces, matches[1])
	}

	return &claims{identifier: identifier, namespaces: namespaces}, nil
}
