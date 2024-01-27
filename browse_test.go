package vlc

import (
	"encoding/json"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestVLC_Browse_Invalid(t *testing.T) {
	t.Parallel()

	testTable := []struct {
		name     string
		browseFn func(vlc *VLC) func(string) (*Browse, error)
		key      string
	}{
		{
			"unable to browse directory path",
			func(vlc *VLC) func(string) (*Browse, error) {
				return vlc.BrowseWithPath
			},
			dirKey,
		},

		{
			"unable to browse directory URI",
			func(vlc *VLC) func(string) (*Browse, error) {
				return vlc.BrowseWithURI
			},
			uriKey,
		},
	}

	for _, testCase := range testTable {
		testCase := testCase

		t.Run(testCase.name, func(t *testing.T) {
			t.Parallel()

			var (
				fetchErr  = errors.New("fetch error")
				directory = "directory source"

				expectedParams = paramMap{
					testCase.key: directory,
				}

				mockClient = &mockClient{
					getFn: func(endpoint string) ([]byte, error) {
						require.Equal(
							t,
							buildQueryEndpoint(baseBrowse, expectedParams),
							endpoint,
						)

						return nil, fetchErr
					},
				}
			)

			vlc := NewVLC(mockClient)

			browse, err := testCase.browseFn(vlc)(directory)

			assert.Nil(t, browse)
			assert.ErrorIs(t, err, fetchErr)
		})
	}
}

func TestVLC_Browse_Valid(t *testing.T) {
	t.Parallel()

	testTable := []struct {
		name     string
		browseFn func(vlc *VLC) func(string) (*Browse, error)
		key      string
	}{
		{
			"successful browse with directory path",
			func(vlc *VLC) func(string) (*Browse, error) {
				return vlc.BrowseWithPath
			},
			dirKey,
		},
		{
			"successful browse with directory uri",
			func(vlc *VLC) func(string) (*Browse, error) {
				return vlc.BrowseWithURI
			},
			uriKey,
		},
	}

	for _, testCase := range testTable {
		testCase := testCase

		t.Run(testCase.name, func(t *testing.T) {
			t.Parallel()

			var (
				directory = "directory source"

				expectedParams = paramMap{
					testCase.key: directory,
				}
				expectedBrowse = &Browse{
					Elements: []File{
						{
							Name: "example",
						},
					},
				}

				mockClient = &mockClient{
					getFn: func(endpoint string) ([]byte, error) {
						require.Equal(
							t,
							buildQueryEndpoint(baseBrowse, expectedParams),
							endpoint,
						)

						return json.Marshal(expectedBrowse)
					},
				}
			)

			vlc := NewVLC(mockClient)

			browseResult, err := testCase.browseFn(vlc)(directory)
			require.NoError(t, err)

			assert.Equal(t, expectedBrowse, browseResult)
		})
	}
}
