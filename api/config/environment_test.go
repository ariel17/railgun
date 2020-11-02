package config

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
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

func TestSetEnvironment(t *testing.T) {
	testCases := []struct{
		name string
		isPresent bool
		value string
	}{
		{"present", true, "value"},
		{"not present", false, "xxx"},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			var err error
			if tc.isPresent {
				err = os.Setenv(envKeyName, tc.value)
			} else {
				err = os.Unsetenv(envKeyName)
			}
			assert.Nil(t, err)
			setEnvironment()
			if tc.isPresent {
				assert.Equal(t, tc.value, environment)
			} else {
				assert.Equal(t, testEnv, environment)
			}
		})
	}
}
