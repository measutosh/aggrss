package auth

import (
	"errors"
	"net/http"
	"strings"
)

// this function will fetch and return the apikey
// this is used in the getUser function in handler-user.go

// it should look something like this
// Authorization: ApiKey {Actual API key}
func GetAPIKey(headers http.Header) (string, error) {
	// pickup the Authorization field value from header using http package
	val := headers.Get("Authorization")
	if val == "" {
		return "", errors.New("no authentication info")
	}

	vals := strings.Split(val, " ")
	if len(vals) != 2 {
		return "", errors.New("malformed auth header")
	}

	if vals[0] != "ApiKey" {
		return "", errors.New("malformed first part of the api key")
	}

	return vals[1], nil
}
