package vlc

import (
	"fmt"
	"regexp"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBuildQueryEndpoint(t *testing.T) {
	t.Parallel()

	t.Run("no params", func(t *testing.T) {
		t.Parallel()

		var (
			baseURL     = "https://example.com"
			expectedURL = baseURL // unchanged
		)

		endpoint := buildQueryEndpoint(baseURL, nil)

		assert.Equal(t, expectedURL, endpoint)
	})

	t.Run("single param", func(t *testing.T) {
		t.Parallel()

		var (
			baseURL = "https://example.com"
			params  = paramMap{
				"key": "value",
			}

			expectedURL = fmt.Sprintf("%s?key=value", baseURL)
		)

		endpoint := buildQueryEndpoint(baseURL, params)

		assert.Equal(t, expectedURL, endpoint)
	})

	t.Run("multiple params", func(t *testing.T) {
		t.Parallel()

		var (
			baseURL = "https://example.com"
			params  = paramMap{
				"key1": "value1",
				"key2": "value2",
			}
		)

		re := regexp.MustCompile(`^https://example\.com\?[A-Za-z0-9]+=[A-Za-z0-9]+&[A-Za-z0-9]+=[A-Za-z0-9]+$`)

		endpoint := buildQueryEndpoint(baseURL, params)

		assert.True(t, re.MatchString(endpoint))

		for key, value := range params {
			assert.Contains(t, endpoint, fmt.Sprintf("%s=%s", key, value))
		}
	})

	t.Run("spaces in params / values", func(t *testing.T) {
		t.Parallel()

		var (
			baseURL = "https://example.com"
			params  = paramMap{
				"a key": "a value",
			}

			expectedURL = fmt.Sprintf("%s?a%%20key=a%%20value", baseURL)
		)

		endpoint := buildQueryEndpoint(baseURL, params)

		assert.Equal(t, expectedURL, endpoint)
	})
}
