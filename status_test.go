package vlc

import (
	"encoding/json"
	"errors"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestVLC_GetStatus(t *testing.T) {
	t.Parallel()

	t.Run("unable to fetch status", func(t *testing.T) {
		t.Parallel()

		var (
			fetchErr   = errors.New("fetch error")
			mockClient = &mockClient{
				getFn: func(endpoint string) ([]byte, error) {
					require.Equal(t, baseStatus, endpoint)

					return nil, fetchErr
				},
			}
		)

		vlc := NewVLC(mockClient)

		status, err := vlc.GetStatus()

		assert.Nil(t, status)
		assert.ErrorIs(t, err, fetchErr)
	})

	t.Run("valid status fetched", func(t *testing.T) {
		t.Parallel()

		var (
			expectedStatus = &Status{
				Version: "random version",
				State:   "playing",
			}

			mockClient = &mockClient{
				getFn: func(endpoint string) ([]byte, error) {
					require.Equal(t, baseStatus, endpoint)

					return json.Marshal(expectedStatus)
				},
			}
		)

		vlc := NewVLC(mockClient)

		status, err := vlc.GetStatus()
		require.NoError(t, err)

		assert.Equal(t, expectedStatus, status)
	})
}

func TestVLC_EmptyPlaylist(t *testing.T) {
	t.Parallel()

	var (
		expectedStatus = &Status{
			Version: "random version",
			State:   "playing",
		}

		expectedParams = paramMap{
			commandKey: emptyCommand,
		}

		mockClient = &mockClient{
			getFn: func(endpoint string) ([]byte, error) {
				require.Equal(
					t,
					buildQueryEndpoint(baseStatus, expectedParams),
					endpoint,
				)

				return json.Marshal(expectedStatus)
			},
		}
	)

	vlc := NewVLC(mockClient)

	status, err := vlc.EmptyPlaylist()
	require.NoError(t, err)

	assert.Equal(t, expectedStatus, status)
}

func TestVLC_PlaySource(t *testing.T) {
	t.Parallel()

	t.Run("invalid play option", func(t *testing.T) {
		t.Parallel()

		var (
			source = "source"
			option = "invalid play option"
		)

		vlc := NewVLC(&mockClient{})

		status, err := vlc.PlaySource(source, option)
		require.Nil(t, status)

		assert.ErrorIs(t, err, errInvalidPlayOption)
	})

	t.Run("no play option", func(t *testing.T) {
		t.Parallel()

		var (
			expectedStatus = &Status{
				Version: "random version",
				State:   "playing",
			}

			source = "source"

			expectedParams = paramMap{
				commandKey: inPlayCommand,
				inputKey:   source,
			}

			mockClient = &mockClient{
				getFn: func(endpoint string) ([]byte, error) {
					require.Equal(
						t,
						buildQueryEndpoint(baseStatus, expectedParams),
						endpoint,
					)

					return json.Marshal(expectedStatus)
				},
			}
		)

		vlc := NewVLC(mockClient)

		status, err := vlc.PlaySource(source)
		require.NoError(t, err)

		assert.Equal(t, expectedStatus, status)
	})

	t.Run("valid play option", func(t *testing.T) {
		t.Parallel()

		var (
			expectedStatus = &Status{
				Version: "random version",
				State:   "playing",
			}

			source = "source"
			option = playNoVideo

			expectedParams = paramMap{
				commandKey: inPlayCommand,
				inputKey:   source,
				optionKey:  option,
			}

			mockClient = &mockClient{
				getFn: func(endpoint string) ([]byte, error) {
					require.Equal(
						t,
						buildQueryEndpoint(baseStatus, expectedParams),
						endpoint,
					)

					return json.Marshal(expectedStatus)
				},
			}
		)

		vlc := NewVLC(mockClient)

		status, err := vlc.PlaySource(source, option)
		require.NoError(t, err)

		assert.Equal(t, expectedStatus, status)
	})
}

func TestVLC_AddToPlaylist(t *testing.T) {
	t.Parallel()

	var (
		expectedStatus = &Status{
			Version: "random version",
			State:   "playing",
		}

		source = "source"

		expectedParams = paramMap{
			commandKey: inEnqueueCommand,
			inputKey:   source,
		}

		mockClient = &mockClient{
			getFn: func(endpoint string) ([]byte, error) {
				require.Equal(
					t,
					buildQueryEndpoint(baseStatus, expectedParams),
					endpoint,
				)

				return json.Marshal(expectedStatus)
			},
		}
	)

	vlc := NewVLC(mockClient)

	status, err := vlc.AddToPlaylist(source)
	require.NoError(t, err)

	assert.Equal(t, expectedStatus, status)
}

func TestVLC_PlayLastActivePlaylistItem(t *testing.T) {
	t.Parallel()

	var (
		expectedStatus = &Status{
			Version: "random version",
			State:   "playing",
		}

		expectedParams = paramMap{
			commandKey: playCommand,
		}

		mockClient = &mockClient{
			getFn: func(endpoint string) ([]byte, error) {
				require.Equal(
					t,
					buildQueryEndpoint(baseStatus, expectedParams),
					endpoint,
				)

				return json.Marshal(expectedStatus)
			},
		}
	)

	vlc := NewVLC(mockClient)

	status, err := vlc.PlayLastActivePlaylistItem()
	require.NoError(t, err)

	assert.Equal(t, expectedStatus, status)
}

func TestVLC_PlayPlaylistItem(t *testing.T) {
	t.Parallel()

	var (
		expectedStatus = &Status{
			Version: "random version",
			State:   "playing",
		}

		id = 10

		expectedParams = paramMap{
			commandKey: playCommand,
			idKey:      strconv.Itoa(id),
		}

		mockClient = &mockClient{
			getFn: func(endpoint string) ([]byte, error) {
				require.Equal(
					t,
					buildQueryEndpoint(baseStatus, expectedParams),
					endpoint,
				)

				return json.Marshal(expectedStatus)
			},
		}
	)

	vlc := NewVLC(mockClient)

	status, err := vlc.PlayPlaylistItem(id)
	require.NoError(t, err)

	assert.Equal(t, expectedStatus, status)
}

func TestVLC_PauseWithLastActivePlaylistItem(t *testing.T) {
	t.Parallel()

	var (
		expectedStatus = &Status{
			Version: "random version",
			State:   "playing",
		}

		expectedParams = paramMap{
			commandKey: pauseCommand,
		}

		mockClient = &mockClient{
			getFn: func(endpoint string) ([]byte, error) {
				require.Equal(
					t,
					buildQueryEndpoint(baseStatus, expectedParams),
					endpoint,
				)

				return json.Marshal(expectedStatus)
			},
		}
	)

	vlc := NewVLC(mockClient)

	status, err := vlc.PauseWithLastActivePlaylistItem()
	require.NoError(t, err)

	assert.Equal(t, expectedStatus, status)
}

func TestVLC_PausePlaylist(t *testing.T) {
	t.Parallel()

	var (
		expectedStatus = &Status{
			Version: "random version",
			State:   "playing",
		}

		id = 5

		expectedParams = paramMap{
			commandKey: pauseCommand,
			idKey:      strconv.Itoa(id),
		}

		mockClient = &mockClient{
			getFn: func(endpoint string) ([]byte, error) {
				require.Equal(
					t,
					buildQueryEndpoint(baseStatus, expectedParams),
					endpoint,
				)

				return json.Marshal(expectedStatus)
			},
		}
	)

	vlc := NewVLC(mockClient)

	status, err := vlc.PausePlaylist(id)
	require.NoError(t, err)

	assert.Equal(t, expectedStatus, status)
}

func TestVLC_ForceResumePlaylist(t *testing.T) {
	t.Parallel()

	var (
		expectedStatus = &Status{
			Version: "random version",
			State:   "playing",
		}

		expectedParams = paramMap{
			commandKey: forceResumeCommand,
		}

		mockClient = &mockClient{
			getFn: func(endpoint string) ([]byte, error) {
				require.Equal(
					t,
					buildQueryEndpoint(baseStatus, expectedParams),
					endpoint,
				)

				return json.Marshal(expectedStatus)
			},
		}
	)

	vlc := NewVLC(mockClient)

	status, err := vlc.ForceResumePlaylist()
	require.NoError(t, err)

	assert.Equal(t, expectedStatus, status)
}

func TestVLC_ForcePausePlaylist(t *testing.T) {
	t.Parallel()

	var (
		expectedStatus = &Status{
			Version: "random version",
			State:   "playing",
		}

		expectedParams = paramMap{
			commandKey: forcePauseCommand,
		}

		mockClient = &mockClient{
			getFn: func(endpoint string) ([]byte, error) {
				require.Equal(
					t,
					buildQueryEndpoint(baseStatus, expectedParams),
					endpoint,
				)

				return json.Marshal(expectedStatus)
			},
		}
	)

	vlc := NewVLC(mockClient)

	status, err := vlc.ForcePausePlaylist()
	require.NoError(t, err)

	assert.Equal(t, expectedStatus, status)
}

func TestVLC_StopPlaylist(t *testing.T) {
	t.Parallel()

	var (
		expectedStatus = &Status{
			Version: "random version",
			State:   "playing",
		}

		expectedParams = paramMap{
			commandKey: stopCommand,
		}

		mockClient = &mockClient{
			getFn: func(endpoint string) ([]byte, error) {
				require.Equal(
					t,
					buildQueryEndpoint(baseStatus, expectedParams),
					endpoint,
				)

				return json.Marshal(expectedStatus)
			},
		}
	)

	vlc := NewVLC(mockClient)

	status, err := vlc.StopPlaylist()
	require.NoError(t, err)

	assert.Equal(t, expectedStatus, status)
}
