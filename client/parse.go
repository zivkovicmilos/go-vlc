package client

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
)

// ParseJSONResponse parses the JSON response into a specific type
func ParseJSONResponse[T any](rawResponse []byte) (*T, error) {
	var response T

	if err := json.Unmarshal(rawResponse, &response); err != nil {
		return nil, fmt.Errorf("unable to unmarshal JSON, %w", err)
	}

	return &response, nil
}

// ParseXMLResponse parses the XML response into a specific type
func ParseXMLResponse[T any](rawResponse []byte) (*T, error) {
	var response T

	if err := xml.Unmarshal(rawResponse, &response); err != nil {
		return nil, fmt.Errorf("unable to unmarshal XML, %w", err)
	}

	return &response, nil
}
