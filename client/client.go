package client

// Client is the remote VLC web server client abstraction
type Client interface {
	// Get executes a GET request, and returns the response body
	Get(endpoint string) ([]byte, error)
}
