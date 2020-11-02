package config

import "os"

const (
	envKeyName = "ENVIRONMENT"
	productionEnv = "production"
	testEnv = "test"
)

var (
	environment string
)

// IsProduction confirms if the current environment is production.
func IsProduction() bool {
	return environment == productionEnv
}

func setEnvironment() {
	var exists bool
	environment, exists = os.LookupEnv(envKeyName)
	if !exists {
		environment = testEnv
	}
}

func init() {
	setEnvironment()
}