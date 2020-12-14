package presenters

import "github.com/ariel17/railgun/api/entities"

// NewDomain is the data used to create a new entry in domain entity.
type NewDomain struct {
	URL string `json:"url"`
}

func (nd NewDomain) ToDomain() *entities.Domain {
	return &entities.Domain{
		URL: nd.URL,
	}
}