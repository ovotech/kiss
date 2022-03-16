package client

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

func LoadToken(tokenPath *string) (*string, error) {
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
