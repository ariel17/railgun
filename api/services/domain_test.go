package services

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/ariel17/railgun/config"
)

func TestGenerateValidationCode(t *testing.T) {
	code := GenerateValidationCode()
	assert.NotEmpty(t, code)
	assert.Equal(t, config.CodeLength, len(code))
}