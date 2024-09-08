package auth 

import (
	"errors"
	"net/http"
	"strings"
)

// extracts an API Key from
// the headers of an HTTP req
// Example:
// Authorization: ApiKey {insert apikey herer}
func GetApiKey(headers http.Header) (string, error) {
	val := headers.Get("Authorization")
	if val == "" {
		return "", errors.New("you must be authenticated to access this endpoint")
	}

	vals := strings.Split(val, " ")
	if len(vals) != 2 {
		return "", errors.New("malformed auth header")
	}
	if vals[0] != "ApiKey" {
		return "", errors.New("malformed first part of auth header")
	}
	return vals[1], nil

}
