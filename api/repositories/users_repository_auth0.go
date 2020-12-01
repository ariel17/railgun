package repositories

import (
	"github.com/ariel17/railgun/api/config"
	"github.com/ariel17/railgun/api/entities"
	"gopkg.in/auth0.v5/management"
)

type Auth0Management management.Management

func (am *Auth0Management) GetUserManager() *management.UserManager {
	return am.User
}

type Auth0UserManagement interface {
	New(domain string, options ...management.ManagementOption) (*management.Management, error)
	WithClientCredentials(clientID, clientSecret string) management.ManagementOption
	GetUserManager() *management.UserManager
}

type auth0Repository struct {
	// TODO replace with custom interface
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
