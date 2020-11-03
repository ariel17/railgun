package repositories

import "github.com/ariel17/railgun/api/entities"

type MockRepository struct {
	User *entities.User
	Err  error
}

func (m *MockRepository) GetByID(_ string) (*entities.User, error) {
	return m.User, m.Err
}

func (m *MockRepository) DeleteByID(_ string) error {
	return m.Err
}

func newUsersRepositoryMock() UsersRepository {
	return &MockRepository{}
}
