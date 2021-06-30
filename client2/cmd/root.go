package cmd

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path"
	"time"

	"github.com/hashicorp/logutils"
	"github.com/spf13/cobra"
)

const (
	defaultTokenPath = ".kube/cache/oidc-login"
)

var (
	secure     bool
	serverAddr string
	timeout    time.Duration
	tokenPath  string
	namespace  string
	debug      bool
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "testcli",
	Short: "Kaluza Infrastructure Secret Service (KISS)",
	Long:  `The KISS CLI allows Kaluza teams to manage secrets in their KMI namespace.`,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	ctx := context.Background()
	cobra.CheckErr(rootCmd.ExecuteContext(ctx))
}

func init() {
	// Persistent flags affect all child subcommands
	rootCmd.PersistentFlags().BoolVar(&secure, "secure", true, "Connection uses TLS if true, else plain TCP")
	rootCmd.PersistentFlags().StringVar(
		&serverAddr,
		"server-addr",
		"localhost:10000",
		"The monitor server address in the format of host:port",
	)
	rootCmd.PersistentFlags().DurationVar(
		&timeout, "timeout", 30*time.Second, "The interval before a connection times out",
	)
	rootCmd.PersistentFlags().StringVar(
		&tokenPath,
		"token-path",
		"",
		"The path to the OIDC JWT used for auth/z. If empty will default to the first file in $HOME/.kube/cache/oidc-login/",
	)
	rootCmd.PersistentFlags().StringVar(&namespace, "namespace", "", "The namespace the secret is scoped to (Required).")
	rootCmd.MarkPersistentFlagRequired("namespace")
	rootCmd.PersistentFlags().BoolVar(&debug, "debug", false, "Enable debug log")

}

func initToken() *string {
	if tokenPath == "" {
		// User hasn't given us a path to a token
		// We'll try to find it ourselves.
		var err error
		tokenPath, err = guessTokenPath()
		if err != nil {
			log.Fatalf("[ERROR] Failed to guess token path, use '-token-path' instead: %s", err)
		}
		log.Printf("[DEBUG] Token not explicitly given, we'll use this one: '%s'", tokenPath)
	}

	token, err := loadToken(&tokenPath)
	if err != nil {
		log.Fatalf("[ERROR] Failed to load token from %s", tokenPath)
	}
	return token
}

func guessTokenPath() (string, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}

	oidcPath := path.Join(homeDir, defaultTokenPath)
	files, err := ioutil.ReadDir(oidcPath)
	if err != nil {
		return "", err
	}

	if len(files) < 1 {
		return "", fmt.Errorf("no token file in %s", oidcPath)
	}

	tokenPath := path.Join(oidcPath, files[0].Name())

	return tokenPath, nil
}

// Initialize logging
func initLogging() {
	logLevel := "WARN"
	if debug {
		logLevel = "DEBUG"
	}
	filter := &logutils.LevelFilter{
		Levels:   []logutils.LogLevel{"DEBUG", "WARN", "ERROR"},
		MinLevel: logutils.LogLevel(logLevel),
		Writer:   os.Stdout,
	}
	log.SetOutput(filter)
}

func loadToken(tokenPath *string) (*string, error) {
	jsonFile, err := os.Open(*tokenPath)
	if err != nil {
		return nil, err
	}
	defer jsonFile.Close()

	bytes, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		return nil, err
	}

	var tokenMap map[string]string
	err = json.Unmarshal([]byte(bytes), &tokenMap)
	if err != nil {
		return nil, err
	}

	idToken, ok := tokenMap["id_token"]
	if !ok {
		return nil, fmt.Errorf("no 'id_token' field in token file")
	}

	return &idToken, nil
}
