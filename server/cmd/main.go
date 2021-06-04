package main

import (
	"flag"
	"os"

	"github.com/rs/zerolog/log"

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
)

func main() {
	flag.Parse()

	if *jwksURL == "" {
		log.Fatal().Msg("-jwks-url is required, see help for more information")
	}

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
