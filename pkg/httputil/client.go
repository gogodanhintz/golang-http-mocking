package httputil

import (
	"io"
	"net/http"
)

type Client interface {
	Do(req *http.Request) (*http.Response, error)
}

// NewHttpClient Create a REAL http client
func NewHttpClient() Client {
	return &http.Client{}
}

// DoGet Helper function to make an HTTP GET request and return the response body
func DoGet(client Client, url string) (string, error) {
	// Create a new Request
	request, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return "", err
	}

	// Send the request via the client (could be real, or could be a mock)
	res, err := client.Do(request)
	if err != nil {
		return "", err
	}

	// Read the response body
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return "", err
	}

	// Return the response body as a string
	return string(body), nil
}
