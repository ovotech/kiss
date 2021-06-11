package client

import (
	"context"
	"log"
	"time"

	pb "github.com/ovotech/kiss/proto"
)

func BindSecret(
	client pb.KISSClient,
	timeout time.Duration,
	namespace, name, serviceAccountName string,
) {
	log.Println("[DEBUG] Creating secret...")
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
		log.Fatalf("[ERROR] Error occurred while binding secret: %v\n", err)
	} else {
		log.Println("[INFO] Successfully bound secret")
	}
}
