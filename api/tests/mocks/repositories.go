package mocks

import (
	"errors"

	"github.com/ariel17/railgun/api/entities"
	"github.com/ariel17/railgun/api/repositories"
	"github.com/ariel17/railgun/api/services"
)

func DomainExists() {
	dr := &repositories.MockDBRepository{}
	services.DomainsRepository = dr
	domain := entities.Domain{
		ID: int64(10),
		UserID: "test-123",
		URL: "ariel17.com.ar",
		Code: "code-123",
		Verified: false,
	}
	dr.Domain = &domain
}

func DomainNotExists() {
	services.DomainsRepository = &repositories.MockDBRepository{}
}

func DomainOperationFails() {
	services.DomainsRepository = &repositories.MockDBRepository{
		Err: errors.New("mocked error :D"),
	}
}
