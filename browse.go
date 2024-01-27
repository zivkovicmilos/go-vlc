package vlc

import (
	"fmt"

	"github.com/zivkovicmilos/go-vlc/client"
)

// executeBrowseRequest executes a GET request and parses the response JSON
func (v *VLC) executeBrowseRequest(params paramMap) (*Browse, error) {
	endpoint := buildQueryEndpoint(baseBrowse, params)

	browseRaw, err := v.client.Get(endpoint)
	if err != nil {
		return nil, fmt.Errorf("unable to execute request, %s, %w", endpoint, err)
	}

	return client.ParseJSONResponse[Browse](browseRaw)
}

// BrowseWithPath browses the given directory file list.
//
// Directory URI is the preferred parameter, so consider using BrowseWithURI, since
// "dir" is deprecated and may be removed in a future release
func (v *VLC) BrowseWithPath(path string) (*Browse, error) {
	params := paramMap{
		dirKey: path,
	}

	return v.executeBrowseRequest(params)
}

// BrowseWithURI browses the given directory URI file list (file://...)
func (v *VLC) BrowseWithURI(uri string) (*Browse, error) {
	params := paramMap{
		uriKey: uri,
	}

	return v.executeBrowseRequest(params)
}
