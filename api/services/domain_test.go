package services

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/ariel17/railgun/api/config"
	"github.com/ariel17/railgun/api/entities"
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

func TestNewDomain(t *testing.T) {
	t.Run("ok", func(t *testing.T) {
		r := &repositories.MockDBRepository{}
		DomainsRepository = r
		d := entities.Domain{
			URL: "blah.com",
		}
		err := NewDomain(&d)
		assert.Nil(t, err)
		assert.Equal(t, int64(10), d.ID)
		assert.Equal(t, "code-123", d.Code)
	})

	t.Run("error", func(t *testing.T) {
		r := &repositories.MockDBRepository{}
		DomainsRepository = r
		r.Err = errors.New("mocked")
		d := entities.Domain{
			URL: "blah.com",
		}
		err := NewDomain(&d)
		assert.Equal(t, int64(0), d.ID)
		assert.NotEqual(t, "", d.Code)
		assert.NotNil(t, err)
	})
}
