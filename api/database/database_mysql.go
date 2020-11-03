package database

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"

	"github.com/ariel17/railgun/api/config"
)

// NewMySQL creates a new instance of MySQL driver.
func NewMySQL() *sql.DB {
	db, err := sql.Open("mysql", dbURL())
	if err != nil {
		panic(err)
	}
	return db
}

func dbURL() string {
	return fmt.Sprintf("%s:%s@(%s:%d)/%s", config.DatabaseUsername, config.DatabasePassword, config.DatabaseHost, config.DatabasePort, config.DatabaseName)
}
