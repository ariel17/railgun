package services

import (
	"math/rand"
	"time"

	"github.com/ariel17/railgun/config"
)

const charset = `abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789[]()*^{}.-,;:_!º"\|?=/&%$<>`

var r *rand.Rand

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

func init() {
	r = rand.New(rand.NewSource(time.Now().UnixNano()))
}
