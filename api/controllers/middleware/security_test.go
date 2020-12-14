package middleware

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
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
		{"valid token", map[string]string{"authorization": "Bearer "+getValidToken()}, true},
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

// curl --request POST \
//  --url https://ariel17.auth0.com/oauth/token \
//  --header 'content-type: application/json' \
//  --data '{"client_id":"BlO4OZFm7BRrXSb807536y75164PwKMc","client_secret":"DXllAOU3ACa2BpEsh-RH3O-VjAJxYosJ5_YD7TLaFLVPfoy_kdoWGvEEVO2tezV7","audience":"https://railgun.ariel17.com.ar","grant_type":"client_credentials"}'
func getValidToken() string {
	type Response struct {
		AccessToken string `json:"access_token"`
		TokenType string `json:"token_type"`
	}
	url := "https://ariel17.auth0.com/oauth/token"
	payload := strings.NewReader("{\"client_id\":\"BlO4OZFm7BRrXSb807536y75164PwKMc\",\"client_secret\":\"DXllAOU3ACa2BpEsh-RH3O-VjAJxYosJ5_YD7TLaFLVPfoy_kdoWGvEEVO2tezV7\",\"audience\":\"https://railgun.ariel17.com.ar\",\"grant_type\":\"client_credentials\"}")
	req, _ := http.NewRequest("POST", url, payload)
	req.Header.Add("content-type", "application/json")
	res, _ := http.DefaultClient.Do(req)
	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)
	var r Response
	_ = json.Unmarshal(body, &r)
	return r.AccessToken
}