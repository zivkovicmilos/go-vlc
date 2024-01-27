package vlc

import (
	"fmt"
	"net/url"

	"github.com/zivkovicmilos/go-vlc/client"
)

const (
	baseVLM        = "requests/vlm.xml"
	baseVLMCommand = "requests/vlm_cmd.xml"
)

// executeVLMRequest executes a GET request and parses the response XML
func (v *VLC) executeVLMRequest(base string, params paramMap) (*VLM, error) {
	endpoint := buildQueryEndpoint(base, params)

	vlmRaw, err := v.client.Get(endpoint)
	if err != nil {
		return nil, fmt.Errorf("unable to execute request, %s, %w", endpoint, err)
	}

	return client.ParseXMLResponse[VLM](vlmRaw)
}

// GetVLMElements fetches the full list of VLM elements
func (v *VLC) GetVLMElements() (*VLM, error) {
	return v.executeVLMRequest(baseVLM, nil)
}

// RunVLMCommand executes the given VLM command
func (v *VLC) RunVLMCommand(command string) (*VLM, error) {
	params := paramMap{
		commandKey: url.QueryEscape(command),
	}

	return v.executeVLMRequest(baseVLMCommand, params)
}
