package main

import (
	"flag"
	"os"

	"github.com/rs/zerolog/log"

	"github.com/ovotech/kiss/pkg/aws"
	"github.com/ovotech/kiss/server"
)

var (
	host          = flag.String("host", "0.0.0.0", "The listening host's interface.")
	port          = flag.Int("port", 10000, "The server's port")
	jwksURL       = flag.String("jwks-url", "", "The URL for the JSON Web Key sets.")
	namespacesKey = flag.String(
		"namespaces-key",
		"cognito:groups",
		"The key containing the list of allowed namespaces in the client's JWT payload.",
	)
	namespacesRegex = flag.String(
		"namespaces-regex",
		"",
		"The optional regex to extract namespaces from the values in '-namespaces-key'. This should match exactly one group. For example, if the namespaces claim list includes 'company-name:default' where 'default' is the namespace allowed, then the regex 'company-name:([1-9a-z-]{1,63})' will successfully extract the namespace.",
	)
	identifierKey = flag.String(
		"identifier-key",
		"email",
		"The key containing the identity of the requester in the client's JWT payload, for auditing purposes.",
	)
	awsWebIdTokenPath = flag.String(
		"token-path",
		"/var/run/secrets/eks.amazonaws.com/serviceaccount/token",
		"Path to the AWS Web Identity Token in the pod. If empty will use default authentication instead (i.e. useful if running locally).",
	)
	iamRoleARN = flag.String(
		"role-arn",
		"",
		"The full ARN of the AWS IAM role used by the controller.",
	)
	awsRegion = flag.String(
		"region",
		"eu-west-1",
		"The AWS region.",
	)
	iamRolePrefix = flag.String(
		"service-account-role-prefix",
		"k8s-sa",
		"Prefix for service account IAM roles to be granted access to secrets. IAM roles that grant access to secrets are named '(prefix_)namespace_service-account-name' where namespace and service-account-name are supplied by clients.",
	)
	secretPrefix = flag.String(
		"secret-prefix",
		"k8s-secret",
		"Prefix for managed AWS Secrets Manager secrets.",
	)
	adminNamespace = flag.String(
		"admin-namespace",
		"",
		"The admin cognito group that can manage secrets in all namespaces",
	)
	roleBindingPrefix = flag.String(
		"rolebinding-prefix",
		"",
		"Prefix for K8s rolebinding name",
	)
	kubeconfigPath = flag.String("kubeconfig-path", "", "Path to Kubeconfig file. Defaults to home directory.")
	enableTracing  = flag.Bool("enable-tracing", true, "Enable DataDog tracing")
)

func main() {
	flag.Parse()

	if *jwksURL == "" {
		log.Fatal().Msg("-jwks-url is required, see help for more information")
	}

	var awsManager *aws.Manager
	if len(*awsWebIdTokenPath) == 0 {
		awsManager = aws.NewManagerWithDefaultConfig(
			*iamRolePrefix,
			*secretPrefix,
			*awsRegion,
			"",
		)
	} else {
		// ARN is required for web id token auth
		if *iamRoleARN == "" {
			log.Fatal().Msgf(
				"Invalid role ARN for controller when using web ID token auth: '%s'. See help for more information.",
				*iamRoleARN,
			)
		}
		awsManager = aws.NewManagerWithWebIdToken(
			*iamRolePrefix,
			*secretPrefix,
			*awsRegion,
			*iamRoleARN,
			*awsWebIdTokenPath,
		)
	}
	server.AWSManager = awsManager

	// Create notification channels
	errChan := make(chan error)
	sigChan := make(chan os.Signal)

	grpcServer, err := server.Run(
		*host,
		*port,
		errChan,
		sigChan,
		jwksURL,
		namespacesKey,
		namespacesRegex,
		identifierKey,
		adminNamespace,
		roleBindingPrefix,
		kubeconfigPath,
		*enableTracing,
	)
	if err != nil {
		log.Fatal().Msgf("Error starting server: %+v", err)
	}

	select {
	case err := <-errChan:
		log.Error().Msgf("Server panic: %+v, connections not terminated gracefully", err)
	case <-sigChan:
		log.Info().Msgf("Received signal to stop, terminating connections gracefully...")
	}

	grpcServer.GracefulStop()
}
