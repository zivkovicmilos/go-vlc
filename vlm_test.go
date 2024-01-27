package vlc

import (
	"encoding/xml"
	"errors"
	"fmt"
	"net/url"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestVLC_GetVLMElements(t *testing.T) {
	t.Parallel()

	t.Run("unable to fetch VLM", func(t *testing.T) {
		t.Parallel()

		var (
			fetchErr   = errors.New("fetch error")
			mockClient = &mockClient{
				getFn: func(endpoint string) ([]byte, error) {
					require.Equal(t, baseVLM, endpoint)

					return nil, fetchErr
				},
			}
		)

		vlc := NewVLC(mockClient)

		vlm, err := vlc.GetVLMElements()

		assert.Nil(t, vlm)
		assert.ErrorIs(t, err, fetchErr)
	})

	t.Run("VLM fetched with no errors", func(t *testing.T) {
		t.Parallel()

		var (
			expectedVLM = &VLM{
				XMLName: xml.Name{
					Local: "vlm",
				},
				Error: "",
			}

			mockClient = &mockClient{
				getFn: func(endpoint string) ([]byte, error) {
					require.Equal(t, baseVLM, endpoint)

					return xml.Marshal(expectedVLM)
				},
			}
		)

		vlc := NewVLC(mockClient)

		vlm, err := vlc.GetVLMElements()
		require.NoError(t, err)

		assert.Equal(t, expectedVLM, vlm)
	})
}

func TestVLC_RunVLMCommand(t *testing.T) {
	t.Parallel()

	t.Run("unable to run VLM command", func(t *testing.T) {
		t.Parallel()

		var (
			command        = "command"
			expectedParams = paramMap{
				commandKey: url.QueryEscape(command),
			}

			fetchErr   = errors.New("fetch error")
			mockClient = &mockClient{
				getFn: func(endpoint string) ([]byte, error) {
					require.Equal(
						t,
						buildQueryEndpoint(baseVLMCommand, expectedParams),
						endpoint,
					)

					return nil, fetchErr
				},
			}
		)

		vlc := NewVLC(mockClient)

		vlm, err := vlc.RunVLMCommand(command)

		assert.Nil(t, vlm)
		assert.ErrorIs(t, err, fetchErr)
	})

	t.Run("VLM command ran with no errors", func(t *testing.T) {
		t.Parallel()

		var (
			command        = "example"
			expectedParams = paramMap{
				commandKey: url.QueryEscape(command),
			}

			expectedVLM = &VLM{
				XMLName: xml.Name{
					Local: "vlm",
				},
				Error: "",
			}

			mockClient = &mockClient{
				getFn: func(endpoint string) ([]byte, error) {
					require.Equal(
						t,
						buildQueryEndpoint(baseVLMCommand, expectedParams),
						endpoint,
					)

					return xml.Marshal(expectedVLM)
				},
			}
		)

		vlc := NewVLC(mockClient)

		vlm, err := vlc.RunVLMCommand(command)
		require.NoError(t, err)

		assert.Equal(t, expectedVLM, vlm)
	})

	t.Run("VLM command ran with errors", func(t *testing.T) {
		t.Parallel()

		var (
			command        = "unknown"
			expectedParams = paramMap{
				commandKey: url.QueryEscape(command),
			}

			expectedVLM = &VLM{
				XMLName: xml.Name{
					Local: "vlm",
				},
				Error: fmt.Sprintf("%s: unknown command", command),
			}

			mockClient = &mockClient{
				getFn: func(endpoint string) ([]byte, error) {
					require.Equal(
						t,
						buildQueryEndpoint(baseVLMCommand, expectedParams),
						endpoint,
					)

					return xml.Marshal(expectedVLM)
				},
			}
		)

		vlc := NewVLC(mockClient)

		vlm, err := vlc.RunVLMCommand(command)
		require.NoError(t, err)

		assert.Equal(t, expectedVLM, vlm)
	})
}
