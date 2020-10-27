package middleware

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/auth0-community/go-auth0"
	"github.com/gin-gonic/gin"
	"gopkg.in/square/go-jose.v2"

	"github.com/ariel17/railgun/api/config"
)

var (
	validator      *auth0.JWTValidator
)

// Claims represents the metadata contained in token.
type Claims struct {
	ID            string
}

// ValidateToken TODO
// See: https://github.com/auth0-community/auth0-go#example
func ValidateToken() gin.HandlerFunc {
	return gin.HandlerFunc(func(c *gin.Context) {
		token, err := validator.ValidateRequest(c.Request)
		if err != nil {
			fmt.Println("Token is not valid:", err)
			c.JSON(http.StatusUnauthorized, gin.H{"error": "token is not valid"})
			c.Abort()
			return
		}

		rawClaims := map[string]interface{}{}
		err = validator.Claims(c.Request, token, &rawClaims)
		if err != nil {
			fmt.Println("Invalid claims:", err)
			c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid claims"})
			c.Abort()
			return
		}

		claims := newClaims(rawClaims)
		c.Set("claims", claims)
		c.Next()
	})
}

// GetClaims extract token claims from context.
func GetClaims(c *gin.Context) (*Claims, error) {
	v, exists := c.Get("claims")
	if !exists {
		return nil, errors.New("claims not found")
	}
	return v.(*Claims), nil
}

func newClaims(claims map[string]interface{}) *Claims {
	return &Claims{
		ID: claims["sub"].(string),
	}
}

func newValidator(tenantDomain, audience string) *auth0.JWTValidator {
	domain := "https://" + tenantDomain + "/"
	client := auth0.NewJWKClient(auth0.JWKClientOptions{URI: domain + ".well-known/jwks.json"}, nil)
	configuration := auth0.NewConfiguration(client, []string{audience}, domain, jose.RS256)
	return auth0.NewValidator(configuration, nil)
}

func init() {
	validator = newValidator(config.Auth0Domain, config.Auth0Audience)
}