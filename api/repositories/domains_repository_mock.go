package repositories

import "github.com/ariel17/railgun/api/entities"

type MockDBRepository struct {
	Domain *entities.Domain
	Err error
}

func (m *MockDBRepository) GetByID(_ int64) (*entities.Domain, error) {
	return m.Domain, m.Err
}

func (m *MockDBRepository) Add(_ *entities.Domain) error {
	return m.Err
}

func (m *MockDBRepository) Update(_ *entities.Domain) error {
	return m.Err
}

func (m *MockDBRepository) DeleteByID(_ int64) error {
	return m.Err
}

func newMockDomainsRepository() DomainsRepository {
	return &MockDBRepository{}
}
