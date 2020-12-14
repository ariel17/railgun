package services

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/ariel17/railgun/api/config"
	"github.com/ariel17/railgun/api/repositories"
)

func TestGenerateValidationCode(t *testing.T) {
	code := generateValidationCode()
	assert.NotEmpty(t, code)
	assert.Equal(t, config.CodeLength, len(code))
}

func TestGetDomain(t *testing.T) {
	t.Run("ok but not found", func(t *testing.T) {
		r := &repositories.MockDBRepository{}
		DomainsRepository = r
		d, err := GetDomain("blah.com")
		assert.Nil(t, err)
		assert.Nil(t, d)
	})

	t.Run("error", func(t *testing.T) {
		r := &repositories.MockDBRepository{}
		DomainsRepository = r
		r.Err = errors.New("mocked")
		d, err := GetDomain("blah.com")
		assert.Nil(t, d)
		assert.NotNil(t, err)
	})
}