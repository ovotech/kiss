package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path"
	"time"

	"github.com/hashicorp/logutils"
	"github.com/ovotech/kiss/client"
	pb "github.com/ovotech/kiss/proto"
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

	pingCmd = flag.NewFlagSet("ping", flag.ExitOnError)

	createSecretCmd    = flag.NewFlagSet("create", flag.ExitOnError)
	createSecretName   = createSecretCmd.String("name", "", "The name of the secret.")
	createSecretValue  = createSecretCmd.String("value", "", "The value of the secret.")
	createSecretPolicy = createSecretCmd.Bool(
		"policy",
		false,
		"Create an AWS IAM policy for reading this secret.",
	)

	listSecretsCmd = flag.NewFlagSet("list", flag.ExitOnError)

	bindSecretCmd            = flag.NewFlagSet("bind", flag.ExitOnError)
	bindSecretName           = bindSecretCmd.String("name", "", "The name of the secret.")
	bindSecretServiceAccount = bindSecretCmd.String(
		"service-account",
		"",
		"The k8s service account that requires access to the secret.",
	)

	updateSecretCmd   = flag.NewFlagSet("update", flag.ExitOnError)
	updateSecretName  = updateSecretCmd.String("name", "", "The name of the secret to update.")
	updateSecretValue = updateSecretCmd.String("value", "", "The new value of the secret.")

	deleteSecretCmd    = flag.NewFlagSet("delete", flag.ExitOnError)
	deleteSecretName   = deleteSecretCmd.String("name", "", "The name of the secret.")
	deleteSecretPolicy = deleteSecretCmd.Bool(
		"policy",
		false,
		"Delete the AWS IAM policy for reading this secret.",
	)

	subcommands = map[string]*flag.FlagSet{
		pingCmd.Name():         pingCmd,
		createSecretCmd.Name(): createSecretCmd,
		listSecretsCmd.Name():  listSecretsCmd,
		bindSecretCmd.Name():   bindSecretCmd,
		updateSecretCmd.Name(): updateSecretCmd,
		deleteSecretCmd.Name(): deleteSecretCmd,
	}
)

func main() {
	setupCommonFlags()

	// Parse and validate subcommand and flags
	// The first argument on the command line is the command
	cmd := subcommands[os.Args[1]]
	if cmd == nil {
		log.Fatalf("[ERROR] unknown subcommand '%s', see help for more details.", os.Args[1])
	}

	// Arguments 2 onwards are flags
	cmd.Parse(os.Args[2:])
	// We initialize logging here because we need -debug from flags
	initLogging()
	validateCommonParams()

	token, err := loadToken(&tokenPath)
	if err != nil {
		log.Fatalf("[ERROR] Failed to load token from %s", tokenPath)
	}

	// Setup gRPC connection and get protobuf client
	conn, err := client.GetConnection(secure, serverAddr, timeout, *token)
	if err != nil {
		log.Fatalf("[ERROR] Error establishing connection: %s", err)
	}
	defer conn.Close()

	kissClient := pb.NewKISSClient(conn)

	// Run subcommand
	switch cmd.Name() {
	case "ping":
		client.Ping(kissClient, timeout, namespace)
	case "create":
		if *createSecretName == "" || *createSecretValue == "" {
			log.Fatalf("[ERROR] -name and -value are required, see help for more details.")
		}
		client.CreateSecret(kissClient, timeout, namespace, *createSecretName, *createSecretValue)
		if *createSecretPolicy {
			client.CreateSecretIAMPolicy(kissClient, timeout, namespace, *createSecretName)
		}
	case "list":
		client.ListSecrets(kissClient, timeout, namespace)
	case "bind":
		if *bindSecretName == "" || *bindSecretServiceAccount == "" {
			log.Fatalf("[ERROR] -name and -service-acount are required, see help for more details")
		}
		client.BindSecret(
			kissClient,
			timeout,
			namespace,
			*bindSecretName,
			*bindSecretServiceAccount,
		)
	case "update":
		if *updateSecretName == "" || *updateSecretValue == "" {
			log.Fatalf("[ERROR] -name and -value are required, see help for more details.")
		}
		client.UpdateSecret(kissClient, timeout, namespace, *updateSecretName, *updateSecretValue)
	case "delete":
		if *deleteSecretName == "" {
			log.Fatalf("[ERROR] -name is required, see help for more details.")
		}
		client.DeleteSecret(kissClient, timeout, namespace, *deleteSecretName)
		if *deleteSecretPolicy {
			client.DeleteSecretIAMPolicy(kissClient, timeout, namespace, *deleteSecretName)
		}
	default:
		log.Fatalf("[ERROR] Unknown command")
	}

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

// Set up common flags used by all commands.
func setupCommonFlags() {
	for _, fs := range subcommands {
		fs.BoolVar(&secure, "secure", true, "Connection uses TLS if true, else plain TCP")
		fs.StringVar(
			&serverAddr,
			"server-addr",
			"localhost:10000",
			"The monitor server address in the format of host:port",
		)
		fs.DurationVar(
			&timeout, "timeout", 30*time.Second, "The interval before a connection times out",
		)
		fs.StringVar(
			&tokenPath,
			"token-path",
			"",
			"The path to the OIDC JWT used for auth/z. If empty will default to the first file in $HOME/.kube/cache/oidc-login/",
		)
		fs.StringVar(&namespace, "namespace", "", "The namespace the secret is scoped to.")
		fs.BoolVar(&debug, "debug", false, "Enable debug log")
	}
}

func validateCommonParams() {
	if namespace == "" {
		log.Fatal("[ERROR] The -namespace flag is required for all commands.")
	}

	if tokenPath == "" {
		// User hasn't given us a path to a token
		// We'll try to find it ourselves.
		var err error
		tokenPath, err = guessTokenPath()
		if err != nil {
			log.Fatalf("[ERROR] Failed to guess token path, use '-token-path' instead: %s", err)
		}
		fmt.Printf("Token not explicitly given, we'll use this one: '%s'\n", tokenPath)
	}
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
