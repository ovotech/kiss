package client

import (
	"context"
	"log"
	"time"

	pb "github.com/ovotech/kiss/proto"
)

// Temporary test function to test gRPC authorization.
func CreateSecret(client pb.KISSClient, timeout time.Duration, namespace, name, value string) {
	log.Println("[DEBUG] Creating secret...")
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	_, err := client.CreateSecret(
		ctx,
		&pb.CreateSecretRequest{
			Metadata: &pb.ClientMeta{Namespace: namespace},
			Name:     name,
			Value:    value,
		},
	)
	if err != nil {
		log.Fatalf("[ERROR] Error ocurred while creating secret: %v\n", err)
	} else {
		log.Println("[INFO] Successfully created secret")
	}
}
