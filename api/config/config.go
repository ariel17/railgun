package config

import "os"

const (
	// CodeLength is the fixed length expected in a validation code.
	CodeLength = 15
	productionEnv = "production"
)

var (
	// Environment refers to the server environment name.
	environment string
)

func IsProduction() bool {
	return environment == productionEnv
}

func init() {
	environment = os.Getenv("ENVIRONMENT")
}