package auth0

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"gopkg.in/auth0.v5/management"
)

func TestNewUserManager(t *testing.T) {
	m := NewUserManager()
	_, ok := m.(*management.UserManager)
	assert.True(t, ok)
}
