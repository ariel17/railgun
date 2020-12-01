package repositories

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewUsersRepository(t *testing.T) {
	usersRepositoryImplementation = newUsersRepositoryMock
	r := NewUsersRepository()
	_, ok := r.(*MockUsersRepository)
	assert.True(t, ok)
}