package serviceAdapters

import "github.com/stretchr/testify/mock"

type MockedTokenService struct {
	mock.Mock
}

func (m *MockedTokenService) GetToken(domain string, length int) (string, error) {
	args := m.Called(domain, length)
	return args.String(0), args.Error(1)
}
