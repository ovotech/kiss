package client

import (
	"crypto/tls"
	"flag"
	"time"

	"github.com/rs/zerolog/log"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"

	clientauthz "github.com/ovotech/kiss/pkg/authz/client"
	pb "github.com/ovotech/kiss/proto"
)

// Run starts the client and executes forever until terminated
func Run(
	secure bool,
	serverAddr string,
	timeout time.Duration,
	accessToken string,
	namespace string,
	name string,
) {
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
	)

	// Establish gRPC connection
	log.Info().Msg("Creating gRPC Dial...")
	conn, err := grpc.Dial(serverAddr, opts...)
	if err != nil {
		log.Fatal().Msgf("Failed to dial: %v", err)
	}
	defer conn.Close()

	client := pb.NewKISSClient(conn)
	ping(client, time.Second*5, namespace, name)
}
