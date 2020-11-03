package repositories

import (
	"database/sql/driver"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"

	"github.com/ariel17/railgun/api/database"
)

func TestDatabaseDomainsRepository_GetByID(t *testing.T) {
	id := 10
	db, mock := database.NewMock()
	r := &databaseDomainsRepository{
		DB: db,
	}
	prepareDomainGetByID(mock, id)
	domain, err := r.GetByID(id)
	assert.Nil(t, err)
	assert.NotNil(t, domain)
	assert.Equal(t, id, domain.ID)
}

func TestDatabaseDomainsRepository_GetByUserID(t *testing.T) {
	t.Fatal("implement me")
}

func TestDatabaseDomainsRepository_Update(t *testing.T) {
	t.Fatal("implement me")
}

func TestDatabaseDomainsRepository_DeleteByID(t *testing.T) {
	t.Fatal("implement me")
}

func prepareDomainGetByID(m sqlmock.Sqlmock, id int) {
	rows := sqlmock.NewRows([]string{"user_id", "url", "code", "verified"}).
		AddRow("auth0-1234", "http://ariel17.com.ar", "12345", "false")
	m.ExpectQuery("SELECT user_id, url, code, verified").
		WithArgs([]driver.Value{id}...).
		WillReturnRows(rows)
}
