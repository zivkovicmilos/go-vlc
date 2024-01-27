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

// BrowseDirectory browses the given directory file list.
//
// URI is the preferred parameter, so consider using BrowseURI, since
// "dir" is deprecated and may be removed in a future release
func (v *VLC) BrowseDirectory(directoryPath string) (*Browse, error) {
	params := paramMap{
		dirKey: directoryPath,
	}

	return v.executeBrowseRequest(params)
}

// BrowseURI browses the given directory URI file list (file://...)
func (v *VLC) BrowseURI(folderURI string) (*Browse, error) {
	params := paramMap{
		uriKey: folderURI,
	}

	return v.executeBrowseRequest(params)
}
