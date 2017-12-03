package db

// ConnectionURL MOCK
type mockConnectionURL struct{}

func (s mockConnectionURL) String() string {
	return "TestConnectinURL"
}
