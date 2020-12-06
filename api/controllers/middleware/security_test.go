package middleware

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestValidateToken(t *testing.T) {
	domain := "ariel17.auth0.com"
	audience := "https://railgun.ariel17.com.ar"
	validator = newValidator(domain, audience)

	testCases := []struct {
		name    string
		headers map[string]string
		isValid bool
	}{
		{"missing token", nil, false},
		{"invalid token", map[string]string{"Authorization": "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNTE2MjM5MDIyfQ.SflKxwRJSMeKKF2QT4fwpMeJf36POk6yJV_adQssw5c"}, false},
		{"valid token", map[string]string{"authorization": "Bearer eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCIsImtpZCI6IlFVTkRORUk1UXpsQk1qYzVOVGN4TXpBNFJVTTNNVEJFTkRKR09FWkNNVFV6TmpZNU1qY3dNdyJ9.eyJpc3MiOiJodHRwczovL2FyaWVsMTcuYXV0aDAuY29tLyIsInN1YiI6IkJsTzRPWkZtN0JSclhTYjgwNzUzNnk3NTE2NFB3S01jQGNsaWVudHMiLCJhdWQiOiJodHRwczovL3JhaWxndW4uYXJpZWwxNy5jb20uYXIiLCJpYXQiOjE2MDcyMTQ2MDgsImV4cCI6MTYwNzMwMTAwOCwiYXpwIjoiQmxPNE9aRm03QlJyWFNiODA3NTM2eTc1MTY0UHdLTWMiLCJndHkiOiJjbGllbnQtY3JlZGVudGlhbHMifQ.EqyMbgmbYjh9OL_9EzV3kDX_6uFro_m4cbVQzJhx7ktSvcH80Zc5ML-idexLk1ZH6Oi8shY6hxVQHnO3JSEN2JE7g_WxCANfjANN7DV4AIEsOoULscKaZrssQ4yYxPmOtZAvSWfM_sC6odbohQmBQTfJ1mUElYDvkUUOED2ncYm5XG-U8cLhZQvGWzUYeAilEdjavCbdhZ470SgtvWrekMvgFiiDVGX4ajMqGlQiHWXYlWBMOSNVwrjTjMjnr6sLumKw3riRXbFl3bgEG2Il5YiTYGgvFY36MdCxxVgD6T_XbytLSy8xTXfhtUYpTifyNvZg2JpJM0BihKjxnhfRrA"}, true},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			response := httptest.NewRecorder()
			gin.SetMode(gin.TestMode)
			c, r := gin.CreateTestContext(response)
			r.Use(ValidateToken())
			r.GET("/test", func(c *gin.Context) {
				c.Status(http.StatusOK)
			})
			c.Request, _ = http.NewRequest(http.MethodGet, "/test", nil)
			if tc.headers != nil {
				for k, v := range tc.headers {
					c.Request.Header.Set(k, v)
				}
			}
			r.ServeHTTP(response, c.Request)
			if tc.isValid {
				assert.Equal(t, http.StatusOK, response.Code)
			} else {
				assert.Equal(t, http.StatusUnauthorized, response.Code)
			}
		})
	}
}

func TestGetClaims(t *testing.T) {
	testCases := []struct{
		name string
		found bool
	}{
		{"found", true},
		{"not found", false},
	}

	for _, tc := range testCases {
		c := &gin.Context{}
		if tc.found {
			c.Set("claims", newClaims(map[string]interface{}{"sub":"auth0-12345"}))
		}
		t.Run(tc.name, func(t *testing.T) {
			claims, err := GetClaims(c)
			if tc.found {
				assert.NotNil(t, claims)
				assert.Nil(t, err)
			} else {
				assert.Nil(t, claims)
				assert.NotNil(t, err)
				assert.Contains(t, err.Error(), "claims not found")
			}
		})
	}
}