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
	// TODO create tests cases

	r := gin.Default()
	r.GET("/domains/:value", GetDomainController)
	ts := httptest.NewServer(r)
	defer ts.Close()

	mocks.SelectExistingDomain()

	rest := resty.New()
	response, err := rest.R().Get(ts.URL + "/domains/ariel17.com.ar")
	assert.Nil(t, err)
	assert.NotNil(t, response)
	assert.Equal(t, http.StatusOK, response.StatusCode())

	body := tests.GetGoldenFile(t, "./testdata/domain_get_ok.json")
	assert.Equal(t, string(body), response.String())
}
