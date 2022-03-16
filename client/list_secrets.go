package client

import (
	"context"
	"fmt"
	"log"
	"time"

	pb "github.com/ovotech/kiss/proto"
)

func ListSecrets(client pb.KISSClient, timeout time.Duration, namespace string) []string {
	log.Println("[DEBUG] Listing secrets...")
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	response, err := client.ListSecrets(
		ctx,
		&pb.ListSecretsRequest{Metadata: &pb.ClientMeta{Namespace: namespace}},
	)

	if err != nil {
		log.Fatalf("[ERROR] Error occurred while listing secrets: %v\n", err)
	} else {
		for _, secret := range response.Secrets {
			fmt.Println(secret)
		}
	}

	return response.Secrets
}
