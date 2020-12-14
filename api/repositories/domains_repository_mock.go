package repositories

import "github.com/ariel17/railgun/api/entities"

type MockDBRepository struct {
	Domain *entities.Domain
	Err error
}

var instance *MockDBRepository

func (m *MockDBRepository) GetByID(_ int64) (*entities.Domain, error) {
	return m.Domain, m.Err
}

func (m *MockDBRepository) GetByURL(_ string) (*entities.Domain, error) {
	return m.Domain, m.Err
}

func (m *MockDBRepository) Add(d *entities.Domain) error {
	if m.Err != nil {
		return m.Err
	}
	d.ID = int64(10)
	d.Code = "code-123"
	return nil
}

func (m *MockDBRepository) Update(_ *entities.Domain) error {
	return m.Err
}

func (m *MockDBRepository) DeleteByID(_ int64) error {
	return m.Err
}