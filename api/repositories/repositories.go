package repositories

import "github.com/ariel17/railgun/api/config"

var (
	isProduction func() bool
)

func init() {
	isProduction = config.IsProduction
}
