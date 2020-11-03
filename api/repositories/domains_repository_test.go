package repositories

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewDomainsRepository(t *testing.T) {
	testCases := []struct {
		name         string
		isProduction bool
		expectedType DomainsRepository
	}{
		{"real", true, &databaseDomainsRepository{}},
		{"mock", false, &mockDBRepository{}},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			isProduction = func() bool {
				return tc.isProduction
			}
			r := NewDomainsRepository()
			assert.IsType(t, tc.expectedType, r)
		})
	}
}

func TestDomainsRepository_GetByID(t *testing.T) {
	t.Fatal("implement me")
}

func TestDomainsRepository_GetByUserID(t *testing.T) {
	t.Fatal("implement me")
}

func TestDomainsRepository_Update(t *testing.T) {
	t.Fatal("implement me")
}

func TestDomainsRepository_DeleteByID(t *testing.T) {
	t.Fatal("implement me")
}
