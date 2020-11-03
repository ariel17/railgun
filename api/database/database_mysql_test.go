package database

import (
	"fmt"
	"testing"

	"github.com/go-playground/assert/v2"

	"github.com/ariel17/railgun/api/config"
)

func TestDBURL(t *testing.T) {
	username := config.DatabaseUsername
	password := config.DatabasePassword
	host := config.DatabaseHost
	port := config.DatabasePort
	name := config.DatabaseName
	defer func() {
		config.DatabaseUsername = username
		config.DatabasePassword = password
		config.DatabaseHost = host
		config.DatabasePort = port
		config.DatabaseName	= name
	}()
	config.DatabaseUsername = "root"
	config.DatabasePassword = "pwd"
	config.DatabaseHost = "1.1.1.1"
	config.DatabasePort = 3306
	config.DatabaseName = "railgun"

	url := fmt.Sprintf("%s:%s@(%s:%d)/%s", config.DatabaseUsername, config.DatabasePassword, config.DatabaseHost, config.DatabasePort, config.DatabaseName)
	assert.Equal(t, url, dbURL())
}
