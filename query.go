package vlc

import (
	"sort"
	"strings"
)

type paramMap map[string]string

// buildQueryEndpoint constructs the query string from the given parameters
func buildQueryEndpoint(baseURL string, params paramMap) string {
	// Check if there are any parameters to add
	if len(params) == 0 {
		return baseURL
	}

	// Sort the map keys, so every
	// query endpoint output is deterministic
	keys := make([]string, 0, len(params))
	for k := range params {
		keys = append(keys, k)
	}

	sort.Strings(keys)

	// Build the query string
	queryParams := make([]string, 0, len(params))

	for _, key := range keys {
		value := params[key]

		// Manually escape the parameters
		escapedKey := strings.ReplaceAll(key, " ", "%20")
		escapedValue := strings.ReplaceAll(value, " ", "%20")

		queryParams = append(queryParams, escapedKey+"="+escapedValue)
	}

	queryString := strings.Join(queryParams, "&")

	// Return the query string prefixed with '?'
	return baseURL + "?" + queryString
}
