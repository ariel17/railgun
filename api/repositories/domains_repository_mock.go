package repositories

import "github.com/ariel17/railgun/api/entities"

type MockDBRepository struct {
	Domain *entities.Domain
	Err error
}

func (m *MockDBRepository) GetByID(_ int) (*entities.Domain, error) {
	return m.Domain, m.Err
}

func newMockDomainsRepository() DomainsRepository {
	return &MockDBRepository{}
}
