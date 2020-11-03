package repositories

type mockDBRepository struct {}

func newMockDomainsRepository() DomainsRepository {
	return &mockDBRepository{}
}
