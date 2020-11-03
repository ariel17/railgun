package repositories

import "github.com/ariel17/railgun/api/entities"

type databaseDomainsRepository struct {
}

func newDatabaseDomainsRepository() DomainsRepository {
	return &databaseDomainsRepository{}
}

func (d *databaseDomainsRepository) GetByID(id int) (*entities.Domain, error) {
	return nil, nil
}
