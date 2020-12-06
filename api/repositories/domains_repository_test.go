package repositories

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewDomainsRepository(t *testing.T) {
	domainsRepositoryImplementation = newMockDomainsRepository
	r := NewDomainsRepository()
	_, ok := r.(*MockDBRepository)
	assert.True(t, ok)
}