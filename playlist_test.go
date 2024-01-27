package vlc

import (
	"encoding/json"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestVLC_GetPlaylist(t *testing.T) {
	t.Parallel()

	t.Run("unable to fetch playlist", func(t *testing.T) {
		t.Parallel()

		var (
			fetchErr   = errors.New("fetch error")
			mockClient = &mockClient{
				getFn: func(endpoint string) ([]byte, error) {
					require.Equal(t, basePlaylist, endpoint)

					return nil, fetchErr
				},
			}
		)

		vlc := NewVLC(mockClient)

		playlist, err := vlc.GetPlaylist()

		assert.Nil(t, playlist)
		assert.ErrorIs(t, err, fetchErr)
	})

	t.Run("valid playlist fetched", func(t *testing.T) {
		t.Parallel()

		var (
			expectedPlaylist = &Playlist{
				Ro:   "ro",
				Type: "leaf",
				Name: "random playlist",
			}

			mockClient = &mockClient{
				getFn: func(endpoint string) ([]byte, error) {
					require.Equal(t, basePlaylist, endpoint)

					return json.Marshal(expectedPlaylist)
				},
			}
		)

		vlc := NewVLC(mockClient)

		playlist, err := vlc.GetPlaylist()
		require.NoError(t, err)

		assert.Equal(t, expectedPlaylist, playlist)
	})
}
