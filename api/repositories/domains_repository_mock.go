package repositories

import "github.com/ariel17/railgun/api/entities"

type mockDBRepository struct {
	Domain *entities.Domain
	Err error
}

func (m *mockDBRepository) GetByID(_ int) (*entities.Domain, error) {
	return m.Domain, m.Err
}

func newMockDomainsRepository() DomainsRepository {
	return &mockDBRepository{}
}
