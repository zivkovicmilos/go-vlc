package http

import (
	"fmt"
	"io"
	"net/http"
)

type RequestAuth struct {
	Username string
	Password string
}

type Client struct {
	baseURL string

	client *http.Client
	auth   RequestAuth
}

// NewClient creates a new instance of the HTTP client
func NewClient(baseURL string, auth RequestAuth) *Client {
	return &Client{
		baseURL: baseURL,
		client:  &http.Client{},
		auth:    auth,
	}
}

func (c *Client) Get(endpoint string) ([]byte, error) {
	// Create the request
	req, err := http.NewRequest(
		http.MethodGet,
		fmt.Sprintf("%s/%s", c.baseURL, endpoint),
		http.NoBody,
	)
	if err != nil {
		return nil, fmt.Errorf("unable to create request, %w", err)
	}

	// Set the basic auth
	req.SetBasicAuth(c.auth.Username, c.auth.Password)

	// Run the request
	response, reqError := c.client.Do(req)
	if reqError != nil {
		return nil, fmt.Errorf("unable to execute request, %w", reqError)
	}

	defer func() {
		_ = response.Body.Close()
	}()

	// Check status code
	statusCode := response.StatusCode
	if !isOKResponse(statusCode) {
		return nil, fmt.Errorf("invalid status code, %d", statusCode)
	}

	responseBody, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, fmt.Errorf("unable to read response body, %w", err)
	}

	return responseBody, nil
}

// isOKResponse validates the response code is valid
func isOKResponse(code int) bool {
	return code >= 200 && code <= 299
}
