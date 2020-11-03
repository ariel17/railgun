package repositories

type databaseDomainsRepository struct {
}

func newDatabaseDomainsRepository() DomainsRepository {
	return &databaseDomainsRepository{}
}
