package database

import (
	"fmt"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/ariel17/railgun/api/config"
)

func TestMain(m *testing.M) {
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

	os.Exit(m.Run())
}

func TestDBURL(t *testing.T) {
	url := fmt.Sprintf("%s:%s@(%s:%d)/%s", config.DatabaseUsername, config.DatabasePassword, config.DatabaseHost, config.DatabasePort, config.DatabaseName)
	assert.Equal(t, url, dbURL())
}

func TestNewMySQL(t *testing.T) {
	t.Run("ok", func(t *testing.T) {
		c := NewMySQL()
		assert.NotNil(t, c)
	})
}