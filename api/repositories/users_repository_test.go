package repositories

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/ariel17/railgun/api/entities"
)

func TestUsersRepository(t *testing.T) {
	usersRepositoryImplementation = newUsersRepositoryMock
	id := "auth0-1234-fake"

	t.Run("get by id", func(t *testing.T) {
		testCases := []struct {
			name       string
			exists     bool
			successful bool
		}{
			{"exists and successful", true, true},
			{"does not exist and successful", false, true},
			{"could exist but fails", true, false},
		}
		for _, tc := range testCases {
			t.Run(tc.name, func(t *testing.T) {
				repository := NewUsersRepository()
				if tc.successful {
					if tc.exists {
						repository.(*MockRepository).User = &entities.User{ID: id}
					}
				} else {
					repository.(*MockRepository).Err = errors.New("mocked error")
				}
				user, err := repository.GetByID(id)
				assert.Equal(t, err == nil, tc.successful)
				assert.Equal(t, user != nil, tc.exists && tc.successful)
			})
		}
	})

	t.Run("delete by id", func(t *testing.T) {
		id := "auth0-1234-fake"
		testCases := []struct {
			name       string
			successful bool
		}{
			{"successful", true},
			{"fails", false},
		}
		for _, tc := range testCases {
			t.Run(tc.name, func(t *testing.T) {
				repository := NewUsersRepository()
				if !tc.successful {
					repository.(*MockRepository).Err = errors.New("mocked error")
				}
				err := repository.DeleteByID(id)
				assert.Equal(t, err == nil, tc.successful)
			})
		}
	})
}
