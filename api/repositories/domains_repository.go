package repositories

import "github.com/ariel17/railgun/api/entities"

// DomainsRepository is the behaviour contract for all Domain's repository
// implementations.
type DomainsRepository interface {
	GetByID(id int) (*entities.Domain, error)
}

var (
	domainsRepositoryImplementation func() DomainsRepository
)

// NewDomainsRepository creates a new instance of Domain's repository
// implementation based on the current environment.
func NewDomainsRepository() DomainsRepository {
	return domainsRepositoryImplementation()
}

func init() {
	domainsRepositoryImplementation = newDatabaseDomainsRepository
}
