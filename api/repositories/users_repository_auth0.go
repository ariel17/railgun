package repositories

import (
	"github.com/ariel17/railgun/api/config"
	"github.com/ariel17/railgun/api/entities"
	"gopkg.in/auth0.v5/management"
)

type auth0Repository struct {
	m *management.Management
}

func newUsersRepositoryAuth0() UsersRepository {
	credentials := management.WithClientCredentials(config.Auth0ManagementClientID, config.Auth0ManagementClientSecret)
	m, err := management.New(config.Auth0Domain, credentials)
	if err != nil {
		panic(err)
	}
	return &auth0Repository{
		m: m,
	}
}

func (a *auth0Repository) GetByID(id string) (*entities.User, error) {
	u, err := a.m.User.Read(id)
	if err != nil {
		return nil, err
	}
	return &entities.User{
		ID: *u.ID,
	}, nil
}

func (a *auth0Repository) DeleteByID(id string) error {
	return a.m.User.Delete(id)
}
