package vlc

import (
	"encoding/json"
	"errors"
	"fmt"
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

func TestVLC_PlayNextInPlaylist(t *testing.T) {
	t.Parallel()

	var (
		expectedStatus = &Status{
			Version: "random version",
			State:   "playing",
		}

		expectedParams = paramMap{
			commandKey: nextCommand,
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

	status, err := vlc.PlayNextInPlaylist()
	require.NoError(t, err)

	assert.Equal(t, expectedStatus, status)
}

func TestVLC_PlayPreviousInPlaylist(t *testing.T) {
	t.Parallel()

	var (
		expectedStatus = &Status{
			Version: "random version",
			State:   "playing",
		}

		expectedParams = paramMap{
			commandKey: previousCommand,
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

	status, err := vlc.PlayPreviousInPlaylist()
	require.NoError(t, err)

	assert.Equal(t, expectedStatus, status)
}

func TestVLC_DeleteFromPlaylist(t *testing.T) {
	t.Parallel()

	var (
		expectedStatus = &Status{
			Version: "random version",
			State:   "playing",
		}

		id = 5

		expectedParams = paramMap{
			commandKey: deleteCommand,
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

	status, err := vlc.DeleteFromPlaylist(id)
	require.NoError(t, err)

	assert.Equal(t, expectedStatus, status)
}

func TestVLC_SortPlaylist(t *testing.T) {
	t.Parallel()

	t.Run("invalid sort ID", func(t *testing.T) {
		t.Parallel()

		var (
			id   = 10 // invalid
			mode = 0  // sort by ID
		)

		vlc := NewVLC(&mockClient{})

		status, err := vlc.SortPlaylist(id, mode)
		require.Nil(t, status)
		assert.ErrorIs(t, err, errInvalidSortMode)
	})

	t.Run("valid sort ID", func(t *testing.T) {
		t.Parallel()

		var (
			expectedStatus = &Status{
				Version: "random version",
				State:   "playing",
			}

			id   = 0
			mode = 0

			expectedParams = paramMap{
				commandKey: sortCommand,
				idKey:      strconv.Itoa(id),
				valKey:     strconv.Itoa(mode),
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

		status, err := vlc.SortPlaylist(id, mode)
		require.NoError(t, err)

		assert.Equal(t, expectedStatus, status)
	})
}

func TestVLC_TogglePlaylistLoop(t *testing.T) {
	t.Parallel()

	var (
		expectedStatus = &Status{
			Version: "random version",
			State:   "playing",
		}

		expectedParams = paramMap{
			commandKey: loopCommand,
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

	status, err := vlc.TogglePlaylistLoop()
	require.NoError(t, err)

	assert.Equal(t, expectedStatus, status)
}

func TestVLC_TogglePlaylistRandom(t *testing.T) {
	t.Parallel()

	var (
		expectedStatus = &Status{
			Version: "random version",
			State:   "playing",
		}

		expectedParams = paramMap{
			commandKey: randomCommand,
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

	status, err := vlc.TogglePlaylistRandom()
	require.NoError(t, err)

	assert.Equal(t, expectedStatus, status)
}

func TestVLC_TogglePlaylistRepeat(t *testing.T) {
	t.Parallel()

	var (
		expectedStatus = &Status{
			Version: "random version",
			State:   "playing",
		}

		expectedParams = paramMap{
			commandKey: repeatCommand,
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

	status, err := vlc.TogglePlaylistRepeat()
	require.NoError(t, err)

	assert.Equal(t, expectedStatus, status)
}

func TestVLC_EnableServiceDiscoveryModule(t *testing.T) {
	t.Parallel()

	var (
		expectedStatus = &Status{
			Version: "random version",
			State:   "playing",
		}

		module = "sap"

		expectedParams = paramMap{
			commandKey: serviceDiscoveryCommand,
			valKey:     module,
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

	status, err := vlc.EnableServiceDiscoveryModule(module)
	require.NoError(t, err)

	assert.Equal(t, expectedStatus, status)
}

func TestVLC_ToggleFullscreen(t *testing.T) {
	t.Parallel()

	var (
		expectedStatus = &Status{
			Version: "random version",
			State:   "playing",
		}

		expectedParams = paramMap{
			commandKey: fullscreenCommand,
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

	status, err := vlc.ToggleFullscreen()
	require.NoError(t, err)

	assert.Equal(t, expectedStatus, status)
}

func TestVLC_SetVolume(t *testing.T) {
	t.Parallel()

	t.Run("invalid volume value", func(t *testing.T) {
		t.Parallel()

		value := "invalid volume"

		vlc := NewVLC(&mockClient{})

		status, err := vlc.SetVolume(value)
		require.Nil(t, status)

		assert.ErrorIs(t, err, errInvalidVolumeValue)
	})

	testTable := []struct {
		name   string
		volume string
	}{
		{
			"volume value with +",
			"+10",
		},
		{
			"volume value with -",
			"-10",
		},
		{
			"volume value with %",
			"10%",
		},
		{
			"volume value as number",
			"10",
		},
	}

	for _, testCase := range testTable {
		testCase := testCase

		t.Run(testCase.name, func(t *testing.T) {
			t.Parallel()

			var (
				expectedStatus = &Status{
					Version: "random version",
					State:   "playing",
				}

				expectedParams = paramMap{
					commandKey: volumeCommand,
					valKey:     testCase.volume,
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

			status, err := vlc.SetVolume(testCase.volume)
			require.NoError(t, err)

			assert.Equal(t, expectedStatus, status)
		})
	}
}

func TestVLC_SeekToValue(t *testing.T) {
	t.Parallel()

	t.Run("invalid seek value", func(t *testing.T) {
		t.Parallel()

		value := "random value"

		vlc := NewVLC(&mockClient{})

		status, err := vlc.SeekToValue(value)
		require.Nil(t, status)
		assert.ErrorIs(t, err, errInvalidSeekValue)
	})

	t.Run("valid seek value", func(t *testing.T) {
		t.Parallel()

		var (
			expectedStatus = &Status{
				Version: "random version",
				State:   "playing",
			}

			value = "+1H:2M"

			expectedParams = paramMap{
				commandKey: seekCommand,
				valKey:     value,
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

		status, err := vlc.SeekToValue(value)
		require.NoError(t, err)

		assert.Equal(t, expectedStatus, status)
	})
}

func TestVLC_AddSubtitle(t *testing.T) {
	t.Parallel()

	var (
		expectedStatus = &Status{
			Version: "random version",
			State:   "playing",
		}

		uri = "subtitle uri"

		expectedParams = paramMap{
			commandKey: addSubtitleCommand,
			valKey:     uri,
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

	status, err := vlc.AddSubtitle(uri)
	require.NoError(t, err)

	assert.Equal(t, expectedStatus, status)
}

func TestVLC_SetPreamp(t *testing.T) {
	t.Parallel()

	t.Run("invalid preamp value", func(t *testing.T) {
		t.Parallel()

		testValues := []struct {
			name  string
			value int
		}{
			{
				"value < -20",
				-30,
			},
			{
				"value > 20",
				30,
			},
		}

		for _, testCase := range testValues {
			testCase := testCase

			t.Run(testCase.name, func(t *testing.T) {
				t.Parallel()

				vlc := NewVLC(&mockClient{})

				status, err := vlc.SetPreamp(testCase.value)
				require.Nil(t, status)
				assert.ErrorIs(t, err, errInvalidPreampGainValue)
			})
		}
	})

	t.Run("valid gain value", func(t *testing.T) {
		t.Parallel()

		var (
			expectedStatus = &Status{
				Version: "random version",
				State:   "playing",
			}

			gain = 10

			expectedParams = paramMap{
				commandKey: preampCommand,
				valKey:     strconv.Itoa(gain),
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

		status, err := vlc.SetPreamp(gain)
		require.NoError(t, err)

		assert.Equal(t, expectedStatus, status)
	})
}

func TestVLC_SetEQ(t *testing.T) {
	t.Parallel()

	t.Run("invalid gain value", func(t *testing.T) {
		t.Parallel()

		testValues := []struct {
			name  string
			value int
		}{
			{
				"value < -20",
				-30,
			},
			{
				"value > 20",
				30,
			},
		}

		for _, testCase := range testValues {
			testCase := testCase

			t.Run(testCase.name, func(t *testing.T) {
				t.Parallel()

				vlc := NewVLC(&mockClient{})

				status, err := vlc.SetEQ(10, testCase.value)
				require.Nil(t, status)
				assert.ErrorIs(t, err, errInvalidPreampGainValue)
			})
		}
	})

	t.Run("valid gain value", func(t *testing.T) {
		t.Parallel()

		var (
			expectedStatus = &Status{
				Version: "random version",
				State:   "playing",
			}

			gain = 10
			band = 10

			expectedParams = paramMap{
				commandKey: equalizerCommand,
				bandKey:    strconv.Itoa(band),
				valKey:     strconv.Itoa(gain),
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

		status, err := vlc.SetEQ(band, gain)
		require.NoError(t, err)

		assert.Equal(t, expectedStatus, status)
	})
}

func TestVLC_EnableEQ(t *testing.T) {
	t.Parallel()

	var (
		expectedStatus = &Status{
			Version: "random version",
			State:   "playing",
		}

		expectedParams = paramMap{
			commandKey: enableeqCommand,
			valKey:     "1", // true
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

	status, err := vlc.EnableEQ(true)
	require.NoError(t, err)

	assert.Equal(t, expectedStatus, status)
}

func TestVLC_SetEQPreset(t *testing.T) {
	t.Parallel()

	var (
		expectedStatus = &Status{
			Version: "random version",
			State:   "playing",
		}

		id = 0

		expectedParams = paramMap{
			commandKey: setpresetCommand,
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

	status, err := vlc.SetEQPreset(id)
	require.NoError(t, err)

	assert.Equal(t, expectedStatus, status)
}

func TestVLC_SelectTitle(t *testing.T) {
	t.Parallel()

	var (
		expectedStatus = &Status{
			Version: "random version",
			State:   "playing",
		}

		id = 0

		expectedParams = paramMap{
			commandKey: titleCommand,
			valKey:     strconv.Itoa(id),
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

	status, err := vlc.SelectTitle(id)
	require.NoError(t, err)

	assert.Equal(t, expectedStatus, status)
}

func TestVLC_SelectChapter(t *testing.T) {
	t.Parallel()

	var (
		expectedStatus = &Status{
			Version: "random version",
			State:   "playing",
		}

		id = 0

		expectedParams = paramMap{
			commandKey: chapterCommand,
			valKey:     strconv.Itoa(id),
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

	status, err := vlc.SelectChapter(id)
	require.NoError(t, err)

	assert.Equal(t, expectedStatus, status)
}

func TestVLC_SelectAudioTrack(t *testing.T) {
	t.Parallel()

	var (
		expectedStatus = &Status{
			Version: "random version",
			State:   "playing",
		}

		id = 0

		expectedParams = paramMap{
			commandKey: audioTrackCommand,
			valKey:     strconv.Itoa(id),
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

	status, err := vlc.SelectAudioTrack(id)
	require.NoError(t, err)

	assert.Equal(t, expectedStatus, status)
}

func TestVLC_SelectVideoTrack(t *testing.T) {
	t.Parallel()

	var (
		expectedStatus = &Status{
			Version: "random version",
			State:   "playing",
		}

		id = 0

		expectedParams = paramMap{
			commandKey: videoTrackCommand,
			valKey:     strconv.Itoa(id),
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

	status, err := vlc.SelectVideoTrack(id)
	require.NoError(t, err)

	assert.Equal(t, expectedStatus, status)
}

func TestVLC_SelectSubtitleTrack(t *testing.T) {
	t.Parallel()

	var (
		expectedStatus = &Status{
			Version: "random version",
			State:   "playing",
		}

		id = 0

		expectedParams = paramMap{
			commandKey: subtitleTrackCommand,
			valKey:     strconv.Itoa(id),
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

	status, err := vlc.SelectSubtitleTrack(id)
	require.NoError(t, err)

	assert.Equal(t, expectedStatus, status)
}

func TestVLC_SetAudioDelay(t *testing.T) {
	t.Parallel()

	var (
		expectedStatus = &Status{
			Version: "random version",
			State:   "playing",
		}

		delay = 0.123

		expectedParams = paramMap{
			commandKey: audioDelayCommand,
			valKey:     fmt.Sprintf("%f", delay),
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

	status, err := vlc.SetAudioDelay(delay)
	require.NoError(t, err)

	assert.Equal(t, expectedStatus, status)
}

func TestVLC_SetSubtitleDelay(t *testing.T) {
	t.Parallel()

	var (
		expectedStatus = &Status{
			Version: "random version",
			State:   "playing",
		}

		delay = 0.123

		expectedParams = paramMap{
			commandKey: subtitleDelayCommand,
			valKey:     fmt.Sprintf("%f", delay),
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

	status, err := vlc.SetSubtitleDelay(delay)
	require.NoError(t, err)

	assert.Equal(t, expectedStatus, status)
}

func TestVLC_SetPlaybackRate(t *testing.T) {
	t.Parallel()

	t.Run("invalid playback rate", func(t *testing.T) {
		t.Parallel()

		rate := 0.0

		vlc := NewVLC(&mockClient{})

		status, err := vlc.SetPlaybackRate(rate)
		require.Nil(t, status)

		assert.ErrorIs(t, err, errInvalidPlaybackRate)
	})

	t.Run("valid playback rate", func(t *testing.T) {
		t.Parallel()

		var (
			expectedStatus = &Status{
				Version: "random version",
				State:   "playing",
			}

			rate = 0.123

			expectedParams = paramMap{
				commandKey: rateCommand,
				valKey:     fmt.Sprintf("%f", rate),
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

		status, err := vlc.SetPlaybackRate(rate)
		require.NoError(t, err)

		assert.Equal(t, expectedStatus, status)
	})
}

func TestVLC_SetAspectRatio(t *testing.T) {
	t.Parallel()

	var (
		expectedStatus = &Status{
			Version: "random version",
			State:   "playing",
		}

		ratio = "16:10"

		expectedParams = paramMap{
			commandKey: aspectRatioCommand,
			valKey:     ratio,
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

	status, err := vlc.SetAspectRatio(ratio)
	require.NoError(t, err)

	assert.Equal(t, expectedStatus, status)
}
