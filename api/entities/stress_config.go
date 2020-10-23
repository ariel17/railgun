package entities

import (
	"net/http"
	"time"
)

// StressConfig contains all parameters required to perform an stress attack
// into a validated domain.
type StressConfig struct {
	Domain *Domain `json:"domain"`
	Headers []http.Header `json:"headers"`
	Body string `json:"body"`
	Retries int `json:"retries"`
	Status int `json:"successful_status"`
	Timeout int `json:"timeout"`
	StartAt time.Time `json:"start_at"`
	Duration time.Duration `json:"duration"`
	RPM int64 `json:"rpm"`
}
