package repositories

import "github.com/ariel17/railgun/api/entities"

type mockRepository struct{
	User *entities.User
	Err error
}

func newUsersRepositoryMock() UsersRepository {
	return &mockRepository{}
}

func (m *mockRepository) GetByID(_ string) (*entities.User, error) {
	return m.User, m.Err
}

func (m *mockRepository) DeleteByID(_ string) error {
	return m.Err
}