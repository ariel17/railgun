package repositories

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/ariel17/railgun/api/entities"
)

func TestNewUsersRepository(t *testing.T) {
	testCases := []struct {
		name         string
		isProduction bool
		expectedType UsersRepository
	}{
		{"real", true, &auth0Repository{}},
		{"mock", false, &mockRepository{}},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			isProduction = func() bool {
				return tc.isProduction
			}
			r := NewUsersRepository()
			assert.IsType(t, tc.expectedType, r)
		})
	}
}

func TestUsersRepository_GetByID(t *testing.T) {
	id := "auth0-1234-fake"
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
					repository.(*mockRepository).User = &entities.User{ID: id}
				}
			} else {
				repository.(*mockRepository).Err = errors.New("mocked error")
			}
			user, err := repository.GetByID(id)
			assert.Equal(t, err == nil, tc.successful)
			assert.Equal(t, user != nil, tc.exists && tc.successful)
		})
	}
}

func TestUsersRepository_DeleteByID(t *testing.T) {
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
				repository.(*mockRepository).Err = errors.New("mocked error")
			}
			err := repository.DeleteByID(id)
			assert.Equal(t, err == nil, tc.successful)
		})
	}
}
