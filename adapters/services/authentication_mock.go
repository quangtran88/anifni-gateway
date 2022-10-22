package serviceAdapters

import "github.com/stretchr/testify/mock"

type MockedAuthenticationService struct {
	mock.Mock
}

func (m *MockedAuthenticationService) CheckEmailOTP(code string, email string) (bool, error) {
	args := m.Called(code, email)
	return args.Bool(0), args.Error(1)
}
