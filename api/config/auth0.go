package config

import (
	"os"
)

var (
	// Auth0Domain is your TENANT + auth0.com.
	Auth0Domain string

	// Auth0Audience is the Auth0 token IDENTIFIER:
	// If using an SPA: CLIENT ID.
	// If using an API: AUDIENCE (Identifier).
	Auth0Audience string

	// Auth0ManagementClientID is the client ID for resource management.
	Auth0ManagementClientID string
	// Auth0ManagementClientSecret is the client's secret for authentication on
	// Auth0's management resources.
	Auth0ManagementClientSecret string
)

func init() {
	Auth0Domain = os.Getenv("AUTH0_DOMAIN")
	Auth0Audience = os.Getenv("AUTH0_AUDIENCE")
	Auth0ManagementClientID = os.Getenv("AUTH0_MANAGEMENT_CLIENT_ID")
	Auth0ManagementClientSecret = os.Getenv("AUTH0_MANAGEMENT_CLIENT_SECRET")
}