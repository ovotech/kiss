package client

import (
	"context"
	"fmt"
	"log"
	"time"

	pb "github.com/ovotech/kiss/proto"
)

func CreateSecretIAMPolicy(client pb.KISSClient, timeout time.Duration, namespace, name string) {
	log.Println("[DEBUG] Creating secret IAM policy...")
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	_, err := client.CreateSecretIAMPolicy(
		ctx,
		&pb.CreateSecretIAMPolicyRequest{
			Metadata: &pb.ClientMeta{Namespace: namespace},
			Name:     name,
		},
	)
	if err != nil {
		log.Fatalf("[ERROR] Error occurred while creating secret IAM policy: %v\n", err)
	} else {
		fmt.Println("Successfully created secret IAM policy")
	}
}
