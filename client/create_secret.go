package client

import (
	"context"
	"fmt"
	"log"
	"time"

	pb "github.com/ovotech/kiss/proto"
)

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
		log.Fatalf("[ERROR] Error occurred while creating secret: %v\n", err)
	} else {
		fmt.Println("Successfully created secret")
	}
}
