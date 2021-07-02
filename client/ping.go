package client

import (
	"context"
	"fmt"
	"log"
	"time"

	pb "github.com/ovotech/kiss/proto"
)

// Temporary test function to test gRPC authorization.
func Ping(client pb.KISSClient, timeout time.Duration, namespace string) {
	log.Println("[DEBUG] Sending ping...")
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	_, err := client.Ping(
		ctx,
		&pb.PingRequest{Metadata: &pb.ClientMeta{Namespace: namespace}},
	)
	if err != nil {
		log.Fatalf("[ERROR] Error ocurred while sending ping: %v\n", err)
	} else {
		fmt.Println("Successfully sent ping")
	}
}
