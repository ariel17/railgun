package database

type DB interface {
	Select(dest interface{}, query string, args ...interface{}) error
}

var (
	implementation func() DB
)

// New creates an instance of the database driver based on the environment.
func New() DB {
	return implementation()
}

func init() {
	implementation = newMySQL
}
