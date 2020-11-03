package repositories

import (
	"github.com/ariel17/railgun/api/database"
	"github.com/ariel17/railgun/api/entities"
)

type databaseDomainsRepository struct {
	DB database.DB
}

func newDatabaseDomainsRepository() DomainsRepository {
	return &databaseDomainsRepository{
		DB: database.New(),
	}
}

func (d *databaseDomainsRepository) GetByID(id int) (*entities.Domain, error) {
	domain := []entities.Domain{}
	err := d.DB.Select(&domain, "SELECT * FROM domains WHERE id = ?", id)
	if err != nil {
		return &entities.Domain{}, err
	}
	return &domain[0], nil
}
