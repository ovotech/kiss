package client

import (
	"context"
	"log"
	"time"

	pb "github.com/ovotech/kiss/proto"
)

func UpdateSecret(client pb.KISSClient, timeout time.Duration, namespace, name, value string) {
	log.Println("[DEBUG] Updating secret...")
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	_, err := client.UpdateSecret(
		ctx,
		&pb.UpdateSecretRequest{
			Metadata: &pb.ClientMeta{Namespace: namespace},
			Name:     name,
			Value:    value,
		},
	)
	if err != nil {
		log.Fatalf("[ERROR] Error occurred while updating secret: %v\n", err)
	} else {
		log.Println("[INFO] Successfully updated secret")
	}
}
