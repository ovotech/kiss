package cmd

import (
	"context"
	"log"

	pb "github.com/ovotech/kiss/proto"
	"github.com/spf13/cobra"
)

var (
	name  string
	value string
)

// createCmd represents the create command
var createCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a new secret in your team's namespace",
	Long:  `Create a new secret in your team's namespace`,
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

		log.Println("[DEBUG] Creating secret...")
		ctx, cancel := context.WithTimeout(context.Background(), timeout)
		defer cancel()

		_, err2 := kissClient.CreateSecret(
			ctx,
			&pb.CreateSecretRequest{
				Metadata: &pb.ClientMeta{Namespace: namespace},
				Name:     name,
				Value:    value,
			},
		)
		if err2 != nil {
			log.Fatalf("[ERROR] Error occurred while creating secret: %v\n", err2)
		} else {
			log.Println("[INFO] Successfully created secret")
		}
	},
}

func init() {
	rootCmd.AddCommand(createCmd)

	createCmd.Flags().StringVar(
		&name,
		"name",
		"",
		"The name of the secret (Required).",
	)
	rootCmd.MarkFlagRequired("name")
	createCmd.Flags().StringVar(
		&value,
		"value",
		"",
		"The value of the secret (Required).",
	)
	rootCmd.MarkFlagRequired("value")
}
