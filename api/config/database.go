package config

import (
	"os"
	"strconv"
)

var (
	// DatabaseUsername is the user name for authentication into database.
	DatabaseUsername string
	// DatabasePassword is the password value for authentication into database.
	DatabasePassword string
	// DatabaseHost is the host value where the database resides.
	DatabaseHost string
	// DatabasePort is the port number where the service is listening.
	DatabasePort int
	// DatabaseName is the schema name.
	DatabaseName string
)

func init() {
	DatabaseUsername = os.Getenv("DATABASE_USERNAME")
	DatabasePassword = os.Getenv("DATABASE_PASSWORD")
	DatabaseHost = os.Getenv("DATABASE_HOST")
	port := os.Getenv("DATABASE_PORT")
	DatabasePort, _ = strconv.Atoi(port)
	DatabaseName = os.Getenv("DATABASE_NAME")
}