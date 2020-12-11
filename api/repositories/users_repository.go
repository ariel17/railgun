package repositories

import (
	"github.com/ariel17/railgun/api/entities"
)

// UsersRepository defines the behavior for an user's repository implementation.
type UsersRepository interface {
	// GetByID retrieves the user using its ID, if exists.
	GetByID(id string) (*entities.User, error)
	// DeleteByID removes the user by its ID, if exists.
	DeleteByID(id string) error
}
