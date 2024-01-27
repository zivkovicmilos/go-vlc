package vlc

import (
	"fmt"

	"github.com/zivkovicmilos/go-vlc/client"
)

// executeStatusRequest executes a GET request and parses the response JSON
func (v *VLC) executePlaylistRequest(params paramMap) (*Playlist, error) {
	endpoint := buildQueryEndpoint(basePlaylist, params)

	playlistRaw, err := v.client.Get(endpoint)
	if err != nil {
		return nil, fmt.Errorf("unable to execute request, %s, %w", endpoint, err)
	}

	return client.ParseJSONResponse[Playlist](playlistRaw)
}

// GetPlaylist fetches the current playlist
func (v *VLC) GetPlaylist() (*Playlist, error) {
	return v.executePlaylistRequest(nil)
}
