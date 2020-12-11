package repositories

import "github.com/ariel17/railgun/api/entities"

// DomainsRepository is the behaviour contract for all Domain's repository
// implementations.
type DomainsRepository interface {
	GetByID(id int64) (*entities.Domain, error)
	GetByURL(url string) (*entities.Domain, error)
	Add(domain *entities.Domain) error
	Update(domain *entities.Domain) error
	DeleteByID(id int64) error
}
