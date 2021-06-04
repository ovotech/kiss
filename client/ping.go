package client

import (
	"context"
	"time"

	"github.com/rs/zerolog/log"

	pb "github.com/ovotech/kiss/proto"
)

// Temporary test function to test gRPC authorization.
func ping(client pb.KISSClient, timeout time.Duration, namespace string, name string) {
	log.Info().Msg("Sending ping...")
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	_, err := client.Ping(
		ctx,
		&pb.PingRequest{Metadata: &pb.ClientMeta{Namespace: namespace, Name: name}},
	)
	if err != nil {
		log.Error().Msgf("Error ocurred while sending ping: %v", err)
	} else {
		log.Info().Msg("Successfully sent ping")
	}
}
