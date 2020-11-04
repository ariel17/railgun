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

func (d *databaseDomainsRepository) GetByID(id int) (*entities.Domain, error) {
	rows, err := d.DB.Query("SELECT user_id, url, code, verified FROM domains WHERE id = ?", id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var userID, url, code, verified string
		if err := rows.Scan(&userID, &url, &code, &verified); err != nil {
			return &entities.Domain{}, err
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
