package repositories

import (
	"testing"
)

func TestDomainsRepository(t *testing.T) {
	domainsRepositoryImplementation = newMockDomainsRepository

	t.Run("get by id", func(t *testing.T) {
		t.Fatal("implement me")
	})

	t.Run("get by user id", func(t *testing.T) {
		t.Fatal("implement me")
	})

	t.Run("update", func(t *testing.T) {
		t.Fatal("implement me")
	})

	t.Run("delete by id", func(t *testing.T) {
		t.Fatal("implement me")
	})
}