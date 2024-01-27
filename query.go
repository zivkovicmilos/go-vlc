package vlc

import "strings"

type paramMap map[string]string

// buildQueryEndpoint constructs the query string from the given parameters
func buildQueryEndpoint(baseURL string, params paramMap) string {
	// Check if there are any parameters to add
	if len(params) == 0 {
		return baseURL
	}

	// Build the query string
	queryParams := make([]string, 0, len(params))

	for key, value := range params {
		// Manually escape the parameters
		escapedKey := strings.ReplaceAll(key, " ", "%20")
		escapedValue := strings.ReplaceAll(value, " ", "%20")

		queryParams = append(queryParams, escapedKey+"="+escapedValue)
	}

	queryString := strings.Join(queryParams, "&")

	// Return the query string prefixed with '?'
	return baseURL + "?" + queryString
}
