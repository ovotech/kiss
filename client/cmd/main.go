package main

import (
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path"
	"time"

	"github.com/hashicorp/logutils"
	"github.com/ovotech/kiss/client"
)

const (
	defaultTokenPath = ".kube/cache/oidc-login"
)

var (
	secure     = flag.Bool("secure", true, "Connection uses TLS if true, else plain TCP")
	serverAddr = flag.String(
		"server-addr",
		"localhost:10000",
		"The monitor server address in the format of host:port",
	)
	timeout = flag.Duration(
		"timeout",
		30*time.Second,
		"The interval before a connection times out",
	)
	tokenPath = flag.String(
		"token-path",
		"",
		"The path to the OIDC JWT used for auth/z. If empty will default to the first file in $HOME/.kube/cache/oidc-login/",
	)
	namespace = flag.String(
		"namespace",
		"",
		"The namespace the secret is scoped to.",
	)
	name = flag.String(
		"name",
		"",
		"The name of the secret",
	)
	debug = flag.Bool("debug", false, "Enable debug log")
)

func main() {
	flag.Parse()

	logLevel := "WARN"
	if *debug {
		logLevel = "DEBUG"
	}
	filter := &logutils.LevelFilter{
		Levels:   []logutils.LogLevel{"DEBUG", "WARN", "ERROR"},
		MinLevel: logutils.LogLevel(logLevel),
		Writer:   os.Stdout,
	}
	log.SetOutput(filter)

	if *namespace == "" {
		log.Fatal("[ERROR] The -namespace is a required parameters for all commands.")
	}

	if *tokenPath == "" {
		// User hasn't given us a path to a token
		// We'll try to find it ourselves.
		var err error
		tokenPath, err = guessTokenPath()
		if err != nil {
			log.Fatalf("[ERROR] Failed to guess token path, use '-token-path' instead: %s", err)
		}
	}

	token, err := loadToken(tokenPath)
	if err != nil {
		log.Fatalf("[ERROR] Failed to load token from %s", *tokenPath)
	}

	client.Run(*secure, *serverAddr, *timeout, *token, *namespace, *name)
}

func guessTokenPath() (*string, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return nil, err
	}

	oidcPath := path.Join(homeDir, defaultTokenPath)
	files, err := ioutil.ReadDir(oidcPath)
	if err != nil {
		return nil, err
	}

	if len(files) < 1 {
		return nil, errors.New(fmt.Sprintf("no token file in %s", oidcPath))
	}

	tokenPath := path.Join(oidcPath, files[0].Name())

	return &tokenPath, nil
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
		return nil, errors.New(fmt.Sprintf("no 'id_token' field in token file"))
	}

	return &idToken, nil
}
