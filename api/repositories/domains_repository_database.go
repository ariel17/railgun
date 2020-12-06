package repositories

import (
	"database/sql"
	"strconv"

	"github.com/ariel17/railgun/api/database"
	"github.com/ariel17/railgun/api/entities"
)

type databaseDomainsRepository struct {
	DB *sql.DB
}

func newDatabaseDomainsRepository() DomainsRepository {
	return &databaseDomainsRepository{
		DB: database.NewMySQL(),
	}
}

func (d *databaseDomainsRepository) GetByID(id int64) (*entities.Domain, error) {
	rows, err := d.DB.Query("SELECT user_id, url, code, verified FROM domains WHERE id = ?", id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var userID, url, code, verified string
		if err := rows.Scan(&userID, &url, &code, &verified); err != nil {
			return nil, err
		}
		v, err := strconv.ParseBool(verified)
		if err != nil {
			return nil, err
		}
		return &entities.Domain{
			ID: id,
			UserID: userID,
			URL: url,
			Code: code,
			Verified: v,
		}, nil
	}
	return nil, nil
}

func (d *databaseDomainsRepository) Add(domain *entities.Domain) error {
	inserted, err := d.DB.Exec("INSERT INTO domains (user_id, url, code, verified) VALUES (?, ?, ?, false)", domain.UserID, domain.URL, domain.Code)
	if err != nil {
		return err
	}
	domain.ID, err = inserted.LastInsertId()
	return err
}

func (d *databaseDomainsRepository) Update(domain *entities.Domain) error {
	_, err := d.DB.Exec("UPDATE domains SET url = ?, code = ?, verified = false WHERE id = ?", domain.URL, domain.Code, domain.ID)
	return err
}
