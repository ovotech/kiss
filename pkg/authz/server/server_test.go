package server

import (
	"fmt"
	"io/ioutil"
	"testing"

	"github.com/golang-jwt/jwt"
	"github.com/ovotech/kiss/pkg/keyfunc"
)

var hmacSampleSecret []byte
var jwks *keyfunc.JWKs

func init() {
	// Load sample key data
	if keyData, e := ioutil.ReadFile("../../../test/hmacTestKey"); e == nil {
		hmacSampleSecret = keyData
	} else {
		panic(e)
	}

	jwks, _ = keyfunc.Get("http://foo/bar")
}

// Creates a mock token
func createMockToken(groups []string, email string) *jwt.Token {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"cognito:groups": groups,
		"email":          email,
	})

	tokenString, err := token.SignedString(hmacSampleSecret)
	if err != nil {
		fmt.Println(err)
	}

	token, err = jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
		return hmacSampleSecret, nil
	})
	if err != nil {
		fmt.Println(err)
	}

	return token
}

func TestGetCustomClaims(t *testing.T) {
	var tests = []struct {
		name  string
		token *jwt.Token
		want  claims
	}{
		{
			"test extract namespace from token",
			createMockToken([]string{"kaluza:mock-group"}, "integration.test@kaluza.com"),
			claims{
				namespaces: []string{"mock-group"},
				identifier: "integration.test@kaluza.com",
			},
		},
	}

	for _, tt := range tests {
		s := serverAuthzInterceptor{
			jwks:            jwks,
			namespacesKey:   "cognito:groups",
			namespacesRegex: "kaluza:([1-9a-z-]{1,63})",
			identifierKey:   "email",
		}
		fmt.Println(tt.token.Raw)
		testname := fmt.Sprintf("%s,%s", tt.name, tt.want.namespaces)
		t.Run(testname, func(t *testing.T) {
			ans, _ := s.getCustomClaims(tt.token)
			if ans.namespaces[0] != tt.want.namespaces[0] {
				t.Errorf("got %s, want %s", ans.namespaces[0], tt.want.namespaces[0])
			}
		})
	}
}
