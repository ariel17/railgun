package auth0

import "gopkg.in/auth0.v5/management"

// MockUserManager TODO
type MockUserManager struct {
	User *management.User
	Err error
}

func (m *MockUserManager) Read(id string, opts ...management.RequestOption) (*management.User, error) {
	return m.User, m.Err
}

func (m *MockUserManager) Delete(id string, opts ...management.RequestOption) error {
	return m.Err
}

func newMockUserManager() UserManager {
	return &MockUserManager{}
}
