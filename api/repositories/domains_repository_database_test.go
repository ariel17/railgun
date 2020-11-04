package repositories

import (
	"database/sql/driver"
	"errors"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"

	"github.com/ariel17/railgun/api/database"
)

var (
	columns = []string{"user_id", "url", "code", "verified"}
)

func TestDatabaseDomainsRepository_GetByID(t *testing.T) {
	id := 10
	testCases := []struct {
		name    string
		isError bool
		values  []driver.Value
		found bool
	}{
		{"found ok", false, []driver.Value{"auth0-1234", "http://ariel17.com.ar", "12345", "false"}, true},
		{"not found ok", false, nil, false},
		{"failed", true, nil, false},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			db, mock := database.NewMock()
			r := &databaseDomainsRepository{
				DB: db,
			}
			if tc.isError {
				prepareDomainGetByIDWithError(mock, id)
			} else {
				prepareDomainGetByID(mock, id, tc.values)
			}
			domain, err := r.GetByID(id)
			if tc.isError {
				assert.Nil(t, domain)
				assert.NotNil(t, err)
				assert.Equal(t, "mocked error", err.Error())
			} else {
				assert.Nil(t, err)
				if tc.found {
					assert.NotNil(t, domain)
					assert.Equal(t, id, domain.ID)
				} else {
					assert.Nil(t, domain)
				}
			}
		})
	}
}

func expectDomainQuery(m sqlmock.Sqlmock, id int) *sqlmock.ExpectedQuery {
	return m.ExpectQuery("SELECT user_id, url, code, verified").
		WithArgs([]driver.Value{id}...)
}

func prepareDomainGetByID(m sqlmock.Sqlmock, id int, values []driver.Value) {
	rows := sqlmock.NewRows(columns)
	expectDomainQuery(m, id).WillReturnRows(rows)
	if values != nil {
		rows.AddRow(values...)
	}
}

func prepareDomainGetByIDWithError(m sqlmock.Sqlmock, id int) {
	expectDomainQuery(m, id).WillReturnError(errors.New("mocked error"))
}
