package repositories

import (
	"github.com/ariel17/railgun/api/auth0"
	"github.com/ariel17/railgun/api/entities"
)

type auth0Repository struct {
	m auth0.UserManager
}

func newUsersRepositoryAuth0() UsersRepository {
	return &auth0Repository{
		m: auth0.NewUserManager(),
	}
}

func (a *auth0Repository) GetByID(id string) (*entities.User, error) {
	u, err := a.m.Read(id)
	if err != nil {
		return nil, err
	}
	return &entities.User{
		ID: *u.ID,
	}, nil
}

func (a *auth0Repository) DeleteByID(id string) error {
	return a.m.Delete(id)
}
