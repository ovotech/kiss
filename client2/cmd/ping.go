package cmd

import (
	"context"
	"log"

	pb "github.com/ovotech/kiss/proto"
	"github.com/spf13/cobra"
)

// pingCmd represents the ping command
var pingCmd = &cobra.Command{
	Use:   "ping",
	Short: "Test connection and authentication",
	Long:  `Ping will connect and authenticate with the KISS server using your oidc token.`,
	Run: func(cmd *cobra.Command, args []string) {
		initLogging()

		token := initToken()

		// Setup gRPC connection and get protobuf client
		conn, err := GetConnection(secure, serverAddr, timeout, *token)
		if err != nil {
			log.Fatalf("[ERROR] Error establishing connection: %s", err)
		}
		defer conn.Close()

		kissClient := pb.NewKISSClient(conn)

		log.Println("[DEBUG] Sending ping...")
		ctx, cancel := context.WithTimeout(context.Background(), timeout)
		defer cancel()

		_, err2 := kissClient.Ping(
			ctx,
			&pb.PingRequest{Metadata: &pb.ClientMeta{Namespace: namespace}},
		)
		if err2 != nil {
			log.Fatalf("[ERROR] Error ocurred while sending ping: %v\n", err2)
		} else {
			log.Println("[INFO] Successfully sent ping")
		}
	},
}

func init() {
	rootCmd.AddCommand(pingCmd)
}
