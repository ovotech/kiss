// Adapted from https://github.com/MicahParks/keyfunc
// All credits to original author

package keyfunc

import (
	"encoding/json"
	"errors"
	"net/http"
	"sync"
	"time"
)

var (

	// ErrKIDNotFound indicates that the given key ID was not found in the JWKs.
	ErrKIDNotFound = errors.New("the given key ID was not found in the JWKs")

	// ErrMissingAssets indicates there are required assets missing to create a public key.
	ErrMissingAssets = errors.New("required assets are missing to create a public key")
)

// ErrorHandler is a function signature that consumes an error.
type ErrorHandler func(err error)

// JSONKey represents a raw key inside a JWKs.
type JSONKey struct {
	Curve       string `json:"crv"`
	Exponent    string `json:"e"`
	ID          string `json:"kid"`
	Modulus     string `json:"n"`
	X           string `json:"x"`
	Y           string `json:"y"`
	precomputed interface{}
}

// JWKs represents a JSON Web Key Set.
type JWKs struct {
	Keys                map[string]*JSONKey
	client              *http.Client
	endBackground       chan struct{}
	endOnce             sync.Once
	jwksURL             string
	mux                 sync.RWMutex
	refreshErrorHandler ErrorHandler
	refreshInterval     *time.Duration
	refreshTimeout      *time.Duration
	refreshUnknownKID   bool
}

// rawJWKs represents a JWKs in JSON format.
type rawJWKs struct {
	Keys []JSONKey `json:"keys"`
}

// New creates a new JWKs from a raw JSON message.
func New(jwksBytes json.RawMessage) (jwks *JWKs, err error) {

	// Turn the raw JWKs into the correct Go type.
	var rawKS rawJWKs
	if err = json.Unmarshal(jwksBytes, &rawKS); err != nil {
		return nil, err
	}

	// Iterate through the keys in the raw JWKs. Add them to the JWKs.
	jwks = &JWKs{
		Keys: make(map[string]*JSONKey, len(rawKS.Keys)),
	}
	for _, key := range rawKS.Keys {
		key := key
		jwks.Keys[key.ID] = &key
	}

	return jwks, nil
}

// EndBackground ends the background goroutine to update the JWKs. It can only happen once and is
// only effective if the
// JWKs has a background goroutine refreshing the JWKs keys.
func (j *JWKs) EndBackground() {
	j.endOnce.Do(func() {
		if j.endBackground != nil {
			close(j.endBackground)
		}
	})
}

// getKey gets the JSONKey from the given KID from the JWKs. It may refresh the JWKs if configured
// to.
func (j *JWKs) getKey(kid string) (jsonKey *JSONKey, err error) {

	// Get the JSONKey from the JWKs.
	var ok bool
	j.mux.RLock()
	jsonKey, ok = j.Keys[kid]
	j.mux.RUnlock()

	// Check if the key was present.
	if !ok {

		// Check to see if configured to refresh on unknown kid.
		if j.refreshUnknownKID {

			// Refresh the JWKs.
			if err = j.refresh(); err != nil && j.refreshErrorHandler != nil {
				j.refreshErrorHandler(err)
				err = nil
			}

			// Lock the JWKs for async safe use.
			j.mux.RLock()
			defer j.mux.RUnlock()

			// Check if the JWKs refresh contained the requested key.
			if jsonKey, ok = j.Keys[kid]; ok {
				return jsonKey, nil
			}
		}

		return nil, ErrKIDNotFound
	}

	return jsonKey, nil
}
