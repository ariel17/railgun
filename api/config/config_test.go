package config

import (
	"testing"

	"github.com/go-playground/assert/v2"
)

func TestIsProduction(t *testing.T) {
	old := environment
	defer func() {
		environment = old
	}()

	testCases := []struct{
		name string
		expected bool
	}{
		{productionEnv, true},
		{"test", false},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T){
			environment = tc.name
			assert.Equal(t, tc.expected, IsProduction())
		})
	}
}
