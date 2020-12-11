package services

import (
	"math/rand"
	"time"

	"github.com/ariel17/railgun/api/config"
	"github.com/ariel17/railgun/api/entities"
	"github.com/ariel17/railgun/api/repositories"
)

const charset = `abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789[]()*^{}.-,;:_!º"\|?=/&%$<>`

var (
	r *rand.Rand
	DomainsRepository repositories.DomainsRepository
)

// GenerateValidationCode creates an unique string to check for in a near future
// for domain ownership validation.
// Source: https://www.calhoun.io/creating-random-strings-in-go/
func GenerateValidationCode() string {
	b := make([]byte, config.CodeLength)
	for i := range b {
		b[i] = charset[r.Intn(len(charset))]
	}
	return string(b)
}

// GetDomain returns the domain data contained in storage, if exists.
func GetDomain(domain string) (*entities.Domain, error) {
	return DomainsRepository.GetByURL(domain)
}

func init() {
	r = rand.New(rand.NewSource(time.Now().UnixNano()))
	DomainsRepository = repositories.NewDatabaseDomainsRepository()
}
