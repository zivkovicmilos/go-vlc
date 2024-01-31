package http

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// newTestServer creates a new test server instance
func newTestServer(t *testing.T, handler http.Handler) *httptest.Server {
	t.Helper()

	mockServer := httptest.NewServer(handler)

	t.Cleanup(func() {
		mockServer.Close()
	})

	return mockServer
}

func TestHTTP_NewClient(t *testing.T) {
	t.Parallel()

	baseURL := "http://example.com"

	auth := RequestAuth{"user", "pass"}

	// Create the client
	client := NewClient(baseURL, auth)

	// Make sure the client is initialized correctly
	assert.Equal(t, baseURL, client.baseURL)
	assert.Equal(t, auth, client.auth)
}

func TestClient_Get(t *testing.T) {
	t.Parallel()

	var (
		response = []byte("response")

		handler = http.HandlerFunc(
			func(w http.ResponseWriter, r *http.Request) {
				w.WriteHeader(http.StatusOK)
				_, err := w.Write(response)

				require.NoError(t, err)
			},
		)

		server = newTestServer(t, handler)
	)

	// Create the client
	client := NewClient(server.URL, RequestAuth{"user", "pass"})
	resp, err := client.Get("example")

	require.NoError(t, err)
	assert.Equal(t, response, resp)
}
