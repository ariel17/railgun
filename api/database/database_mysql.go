package database

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"

	"github.com/ariel17/railgun/api/config"
)

type mysqlDB struct {
	DBX *sqlx.DB
}

func (m *mysqlDB) Select(dest interface{}, query string, args ...interface{}) error {
	return m.DBX.Select(dest, query, args...)
}

func newMySQL() DB {
	db, err := sqlx.Connect("mysql", dbURL())
	if err != nil {
		panic(err)
	}
	return &mysqlDB{
		DBX: db,
	}
}

func dbURL() string {
	return fmt.Sprintf("%s:%s@(%s:%d)/%s", config.DatabaseUsername, config.DatabasePassword, config.DatabaseHost, config.DatabasePort, config.DatabaseName)
}
