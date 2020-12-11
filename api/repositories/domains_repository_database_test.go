package repositories

import (
	"database/sql/driver"
	"errors"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"

	"github.com/ariel17/railgun/api/database"
	"github.com/ariel17/railgun/api/entities"
)

var (
	columnsByID  = []string{"user_id", "url", "code", "verified"}
	columnsByURL = []string{"id", "user_id", "url", "code", "verified"}
)

func TestNewDatabaseDomainRepository(t *testing.T) {
	r := NewDatabaseDomainsRepository()
	_, ok := r.(*databaseDomainsRepository)
	assert.True(t, ok)
}

func TestDatabaseDomainsRepository_GetByID(t *testing.T) {
	id := int64(10)
	testCases := []struct {
		name    string
		isError bool
		values  []driver.Value
		found   bool
	}{
		{"found ok", false, []driver.Value{"auth0-1234", "ariel17.com.ar", "12345", "false"}, true},
		{"found failed by not boolean", true, []driver.Value{"auth0-1234", "ariel17.com.ar", "12345", "wat?"}, true},
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

func TestDatabaseDomainsRepository_GetByURL(t *testing.T) {
	url := "ariel17.com.ar"
	testCases := []struct {
		name    string
		isError bool
		values  []driver.Value
		found   bool
	}{
		{"found ok", false, []driver.Value{int64(10), "auth0-1234", "ariel17.com.ar", "12345", "false"}, true},
		{"found failed by not boolean", true, []driver.Value{int64(0), "auth0-1234", "ariel17.com.ar", "12345", "wat?"}, true},
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
				prepareDomainGetByURLWithError(mock, url)
			} else {
				prepareDomainGetByURL(mock, url, tc.values)
			}
			domain, err := r.GetByURL(url)
			if tc.isError {
				assert.Nil(t, domain)
				assert.NotNil(t, err)
				assert.Equal(t, "mocked error", err.Error())
			} else {
				assert.Nil(t, err)
				if tc.found {
					assert.NotNil(t, domain)
					assert.Equal(t, url, domain.URL)
				} else {
					assert.Nil(t, domain)
				}
			}
		})
	}
}

func TestDatabaseDomainsRepository_Add(t *testing.T) {
	id := int64(5)
	testCases := []struct {
		name    string
		isError bool
	}{
		{"ok", false},
		{"failed", true},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			domain := entities.Domain{
				UserID: "fake-123",
				URL:    "ariel17.com.ar",
				Code:   "random-123",
			}
			db, mock := database.NewMock()
			r := &databaseDomainsRepository{
				DB: db,
			}
			if tc.isError {
				prepareDomainInsertWithError(mock, domain)
			} else {
				prepareDomainInsert(mock, domain, id)
			}
			err := r.Add(&domain)
			if tc.isError {
				assert.Equal(t, int64(0), domain.ID)
				assert.NotNil(t, err)
				assert.Equal(t, "mocked error", err.Error())
			} else {
				assert.Equal(t, id, domain.ID)
				assert.Nil(t, err)
			}
		})
	}
}

func TestDatabaseDomainsRepository_Update(t *testing.T) {
	testCases := []struct {
		name    string
		isError bool
	}{
		{"ok", false},
		{"failed", true},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			domain := entities.Domain{
				UserID: "fake-123",
				URL:    "ariel17.com.ar",
				Code:   "random-123",
			}
			db, mock := database.NewMock()
			r := &databaseDomainsRepository{
				DB: db,
			}
			if tc.isError {
				prepareDomainUpdateWithError(mock, domain)
			} else {
				prepareDomainUpdate(mock, domain)
			}
			err := r.Update(&domain)
			if tc.isError {
				assert.NotNil(t, err)
				assert.Equal(t, "mocked error", err.Error())
			} else {
				assert.Nil(t, err)
			}
		})
	}
}

func TestDatabaseDomainsRepository_Delete(t *testing.T) {
	id := int64(10)
	testCases := []struct {
		name    string
		isError bool
	}{
		{"ok", false},
		{"failed", true},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			db, mock := database.NewMock()
			r := &databaseDomainsRepository{
				DB: db,
			}
			if tc.isError {
				prepareDomainDeleteByIDWithError(mock, id)
			} else {
				prepareDomainDeleteByID(mock, id)
			}
			err := r.DeleteByID(id)
			if tc.isError {
				assert.NotNil(t, err)
				assert.Equal(t, "mocked error", err.Error())
			} else {
				assert.Nil(t, err)
			}
		})
	}
}

func expectDomainSelectByID(m sqlmock.Sqlmock, id int64) *sqlmock.ExpectedQuery {
	return m.ExpectQuery("SELECT user_id, url, code, verified").
		WithArgs([]driver.Value{id}...)
}

func expectDomainSelectByURL(m sqlmock.Sqlmock, url string) *sqlmock.ExpectedQuery {
	return m.ExpectQuery("SELECT id, user_id, url, code, verified").
		WithArgs([]driver.Value{url}...)
}

func expectDomainInsert(m sqlmock.Sqlmock, domain entities.Domain) *sqlmock.ExpectedExec {
	return m.ExpectExec("INSERT INTO domain").
		WithArgs(domain.UserID, domain.URL, domain.Code)
}

func expectDomainUpdate(m sqlmock.Sqlmock, domain entities.Domain) *sqlmock.ExpectedExec {
	return m.ExpectExec("UPDATE domains SET").
		WithArgs(domain.URL, domain.Code, domain.ID)
}

func expectDomainDelete(m sqlmock.Sqlmock, id int64) *sqlmock.ExpectedExec {
	return m.ExpectExec("DELETE domains WHERE id").WithArgs(id)
}

func prepareDomainGetByID(m sqlmock.Sqlmock, id int64, values []driver.Value) {
	rows := sqlmock.NewRows(columnsByID)
	expectDomainSelectByID(m, id).WillReturnRows(rows)
	if values != nil {
		rows.AddRow(values...)
	}
}

func prepareDomainGetByIDWithError(m sqlmock.Sqlmock, id int64) {
	expectDomainSelectByID(m, id).WillReturnError(errors.New("mocked error"))
}

func prepareDomainGetByURL(m sqlmock.Sqlmock, url string, values []driver.Value) {
	rows := sqlmock.NewRows(columnsByURL)
	expectDomainSelectByURL(m, url).WillReturnRows(rows)
	if values != nil {
		rows.AddRow(values...)
	}
}

func prepareDomainGetByURLWithError(m sqlmock.Sqlmock, url string) {
	expectDomainSelectByURL(m, url).WillReturnError(errors.New("mocked error"))
}

func prepareDomainInsert(m sqlmock.Sqlmock, domain entities.Domain, id int64) {
	expectDomainInsert(m, domain).WillReturnResult(sqlmock.NewResult(id, 1))
}

func prepareDomainInsertWithError(m sqlmock.Sqlmock, domain entities.Domain) {
	expectDomainInsert(m, domain).WillReturnError(errors.New("mocked error"))
}

func prepareDomainUpdate(m sqlmock.Sqlmock, domain entities.Domain) {
	expectDomainUpdate(m, domain).WillReturnResult(sqlmock.NewResult(domain.ID, 1))
}

func prepareDomainUpdateWithError(m sqlmock.Sqlmock, domain entities.Domain) {
	expectDomainUpdate(m, domain).WillReturnError(errors.New("mocked error"))
}

func prepareDomainDeleteByID(m sqlmock.Sqlmock, id int64) {
	expectDomainDelete(m, id).WillReturnResult(sqlmock.NewResult(1, 1))
}

func prepareDomainDeleteByIDWithError(m sqlmock.Sqlmock, id int64) {
	expectDomainDelete(m, id).WillReturnError(errors.New("mocked error"))
}
