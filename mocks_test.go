package vlc

type getDelegate func(string) ([]byte, error)

type mockClient struct {
	getFn getDelegate
}

func (m *mockClient) Get(endpoint string) ([]byte, error) {
	if m.getFn != nil {
		return m.getFn(endpoint)
	}

	return nil, nil
}
