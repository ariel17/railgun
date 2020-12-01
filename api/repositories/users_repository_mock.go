package repositories

import "github.com/ariel17/railgun/api/entities"

type MockUsersRepository struct {
	User *entities.User
	Err  error
}

func (m *MockUsersRepository) GetByID(_ string) (*entities.User, error) {
	return m.User, m.Err
}

func (m *MockUsersRepository) DeleteByID(_ string) error {
	return m.Err
}

func newUsersRepositoryMock() UsersRepository {
	return &MockUsersRepository{}
}
