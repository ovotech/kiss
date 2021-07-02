package client

import (
	"context"
	"fmt"
	"log"
	"time"

	pb "github.com/ovotech/kiss/proto"
)

func DeleteSecret(client pb.KISSClient, timeout time.Duration, namespace, name string) {
	log.Println("[DEBUG] Deleting secret...")
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	_, err := client.DeleteSecret(
		ctx,
		&pb.DeleteSecretRequest{
			Metadata: &pb.ClientMeta{Namespace: namespace},
			Name:     name,
		},
	)
	if err != nil {
		log.Fatalf("[ERROR] Error occurred while deleting secret: %v\n", err)
	} else {
		fmt.Println("Successfully deleted secret")
	}
}
