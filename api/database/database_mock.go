package database

import (
	"database/sql"

	"github.com/DATA-DOG/go-sqlmock"
)

type MockDB struct {
	mysqlDB
	DB *sql.DB
	Mock sqlmock.Sqlmock
}

func newMock() DB {
	db, mock, err := sqlmock.New()
	if err != nil {
		panic(err)
	}
	return &MockDB{
		DB: db,
		Mock: mock,
	}
}
