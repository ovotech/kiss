package client

import (
	"context"
	"fmt"
	"log"
	"time"

	pb "github.com/ovotech/kiss/proto"
)

func BindSecret(
	client pb.KISSClient,
	timeout time.Duration,
	namespace, name, serviceAccountName string,
) {
	log.Println("[DEBUG] Binding secret to service acccount...")
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	_, err := client.BindSecret(
		ctx,
		&pb.BindSecretRequest{
			Metadata:           &pb.ClientMeta{Namespace: namespace},
			Name:               name,
			ServiceAccountName: serviceAccountName,
		},
	)
	if err != nil {
		log.Fatalf("[ERROR] Error occurred while binding secret to service account: %v\n", err)
	} else {
		fmt.Println("Successfully bound secret to service account")
	}
}
