package repositories

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"gopkg.in/auth0.v5/management"

	"github.com/ariel17/railgun/api/auth0"
)

func TestAuth0Repository_GetByID(t *testing.T) {
	t.Run("ok", func(t *testing.T) {
		id := "still-fake"
		r := NewUsersRepositoryAuth0().(*auth0Repository)
		r.m = &auth0.MockUserManager{
			User: &management.User{
				ID: &id,
			},
			Err: nil,
		}
		u, err := r.GetByID("fake-id")
		assert.Nil(t, err)
		assert.Equal(t, id, u.ID)
	})

	t.Run("failed", func(t *testing.T) {
		r := NewUsersRepositoryAuth0().(*auth0Repository)
		r.m = &auth0.MockUserManager{
			Err: errors.New("mocked error"),
		}
		_, err := r.GetByID("fake-id")
		assert.NotNil(t, err)
	})
}

func TestAuth0Repository_DeleteByID(t *testing.T) {
	t.Run("ok", func(t *testing.T) {
		r := NewUsersRepositoryAuth0().(*auth0Repository)
		r.m = &auth0.MockUserManager{}
		err := r.DeleteByID("fake-id")
		assert.Nil(t, err)
	})

	t.Run("fail", func(t *testing.T) {
		r := NewUsersRepositoryAuth0().(*auth0Repository)
		r.m = &auth0.MockUserManager{
			Err: errors.New("mocked error"),
		}
		err := r.DeleteByID("fake-id")
		assert.NotNil(t, err)
	})
}
