package vlc

import (
	"fmt"

	"github.com/zivkovicmilos/go-vlc/client"
)

// GetPlaylist fetches the current playlist
func (v *VLC) GetPlaylist() (*Playlist, error) {
	playlistRaw, err := v.client.Get(basePlaylist)
	if err != nil {
		return nil, fmt.Errorf("unable to fetch playlist, %w", err)
	}

	return client.ParseJSONResponse[Playlist](playlistRaw)
}
