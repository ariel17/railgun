package repositories

// DomainsRepository is the behaviour contract for all Domain's repository
// implementations.
type DomainsRepository interface {
}

// NewDomainsRepository creates a new instance of Domain's repository
// implementation based on the current environment.
func NewDomainsRepository() DomainsRepository {
	if isProduction() {
		return newDatabaseDomainsRepository()
	}
	return newMockDomainsRepository()
}
