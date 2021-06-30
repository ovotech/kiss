package client

import (
	"context"
	"log"
	"time"

	pb "github.com/ovotech/kiss/proto"
)

func DeleteSecretIAMPolicy(client pb.KISSClient, timeout time.Duration, namespace, name string) {
	log.Println("[DEBUG] Deleting secret IAM policy...")
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	_, err := client.DeleteSecretIAMPolicy(
		ctx,
		&pb.DeleteSecretIAMPolicyRequest{
			Metadata: &pb.ClientMeta{Namespace: namespace},
			Name:     name,
		},
	)
	if err != nil {
		log.Fatalf("[ERROR] Error occurred while deleting secret IAM policy: %v\n", err)
	} else {
		log.Println("[INFO] Successfully deleted secret IAM policy")
	}
}
