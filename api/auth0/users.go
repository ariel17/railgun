package auth0

import (
	"gopkg.in/auth0.v5/management"

	"github.com/ariel17/railgun/api/config"
)

// UserManager TODO
type UserManager interface {
	Read(id string, opts ...management.RequestOption) (*management.User, error)
	Delete(id string, opts ...management.RequestOption) error
}

var implementation func() UserManager

// NewUserManager TODO
func NewUserManager() UserManager {
	return implementation()
}

func newAuth0UserManager() UserManager {
	credentials := management.WithClientCredentials(config.Auth0ManagementClientID, config.Auth0ManagementClientSecret)
	m, err := management.New(config.Auth0Domain, credentials)
	if err != nil {
		panic(err)
	}
	return m.User
}

func init() {
	implementation = newAuth0UserManager
}