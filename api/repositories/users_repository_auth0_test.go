package repositories

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"gopkg.in/auth0.v5/management"

	"github.com/ariel17/railgun/api/auth0"
)

func TestAuth0Repository_GetByID(t *testing.T) {
	id := "still-fake"
	r := newUsersRepositoryAuth0().(*auth0Repository)
	r.m = &auth0.MockUserManager{
		User: &management.User{
			ID: &id,
		},
		Err: nil,
	}
	u, err := r.GetByID("fake-id")
	assert.Nil(t, err)
	assert.Equal(t, id, u.ID)
}

func TestAuth0Repository_DeleteByID(t *testing.T) {
	r := newUsersRepositoryAuth0().(*auth0Repository)
	r.m = &auth0.MockUserManager{}
	err := r.DeleteByID("fake-id")
	assert.Nil(t, err)
}