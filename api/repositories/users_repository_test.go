package repositories

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewUsersRepository(t *testing.T) {
	r := newUsersRepositoryMock()
	_, ok := r.(*MockUsersRepository)
	assert.True(t, ok)
}
