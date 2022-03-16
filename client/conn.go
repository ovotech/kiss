package client

import (
	"context"
	"crypto/tls"
	"flag"
	"log"
	"strconv"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"

	clientauthz "github.com/ovotech/kiss/pkg/authz/client"
)

// Run starts the client and executes forever until terminated
func GetConnection(
	secure bool,
	serverAddr string,
	timeout time.Duration,
	accessToken string,
) (*grpc.ClientConn, error) {
	flag.Parse()

	authInterceptor := clientauthz.NewClientAuthInterceptor(accessToken)

	// Prepare gRPC dial options
	var opts []grpc.DialOption
	if secure {
		creds := credentials.NewTLS(&tls.Config{
			MinVersion: tls.VersionTLS12,
		})
		opts = append(opts, grpc.WithTransportCredentials(creds))
	} else {
		opts = append(opts, grpc.WithInsecure())
	}
	opts = append(opts,
		grpc.WithBlock(),
		grpc.WithUnaryInterceptor(authInterceptor.Unary()),
		grpc.FailOnNonTempDialError(true),
	)

	// Establish gRPC connection
	log.Printf("[DEBUG] Creating gRPC Dial... (%s) secure=%s", serverAddr, strconv.FormatBool(secure))
	// conn, err := grpc.Dial(serverAddr, opts...)
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	conn, err := grpc.DialContext(ctx, serverAddr, opts...)

	return conn, err
}
