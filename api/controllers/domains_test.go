package controllers

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/go-resty/resty/v2"
	"github.com/stretchr/testify/assert"

	"github.com/ariel17/railgun/api/tests"
	"github.com/ariel17/railgun/api/tests/mocks"
)

func TestGetDomainController(t *testing.T) {
	r := gin.Default()
	r.GET("/domains/:value", GetDomainController)
	ts := httptest.NewServer(r)
	defer ts.Close()

	rest := resty.New()

	testCases := []struct {
		name       string
		scenario   func()
		status     int
		goldenPath string
	}{
		{"found", mocks.DomainExists, http.StatusOK, "./testdata/domain_get_ok.json"},
		{"not found", mocks.DomainNotExists, http.StatusNotFound, ""},
		{"fails by error", mocks.DomainOperationFails, http.StatusInternalServerError, "./testdata/domain_get_error.json"},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			tc.scenario()
			response, err := rest.R().Get(ts.URL + "/domains/ariel17.com.ar")
			assert.Nil(t, err)
			assert.NotNil(t, response)
			assert.Equal(t, tc.status, response.StatusCode())

			if tc.goldenPath != "" {
				body := tests.GetGoldenFile(t, tc.goldenPath)
				assert.Equal(t, string(body), response.String())
			}
		})
	}
}

func TestNewDomainController(t *testing.T) {
	r := gin.Default()
	r.POST("/domains", NewDomainController)
	ts := httptest.NewServer(r)
	defer ts.Close()

	rest := resty.New()

	testCases := []struct {
		name       string
		scenario   func()
		status     int
		goldenPathInput string
		goldenPathOutput string
	}{
		{"new ok", mocks.DomainExists, http.StatusCreated, "./testdata/domain_post_ok_input.json", "./testdata/domain_post_ok_output.json"},
		{"fails by invalid body", mocks.DomainNotExists, http.StatusBadRequest, "./testdata/domain_post_400_input.json", "./testdata/domain_post_400_output.json"},
		{"fails by error", mocks.DomainOperationFails, http.StatusInternalServerError, "./testdata/domain_post_error_input.json", "./testdata/domain_post_error_output.json"},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			tc.scenario()

			input := tests.GetGoldenFile(t, tc.goldenPathInput)
			response, err := rest.R().
				SetHeader("Content-Type", "application/json").
				SetBody(input).Post(ts.URL + "/domains")
			assert.Nil(t, err)
			assert.NotNil(t, response)
			assert.Equal(t, tc.status, response.StatusCode())

			if tc.goldenPathOutput != "" {
				body := tests.GetGoldenFile(t, tc.goldenPathOutput)
				assert.Equal(t, string(body), response.String())
			}
		})
	}

}
