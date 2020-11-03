package database

import (
	"github.com/ariel17/railgun/api/config"
)

type DB interface {
	Select(dest interface{}, query string, args ...interface{}) error
}

var (
	isProduction func() bool
)

// New creates an instance of the database driver based on the environment.
func New() DB {
	if isProduction() {
		return newMySQL()
	}
	return newMock()
}

func init() {
	isProduction = config.IsProduction
}
