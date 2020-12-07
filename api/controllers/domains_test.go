package controllers

import (
	"net/http"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/go-resty/resty/v2"
	"github.com/stretchr/testify/assert"

	"github.com/ariel17/railgun/api/tests"
)

func TestGetDomainController(t *testing.T) {
	// TODO create tests cases

	r := gin.Default()
	r.GET("/domain/:url", GetDomainController)
	server := tests.CreateAndStartTestServer(r)
	defer tests.StopTestServer(server)

	rest := resty.New()
	response, err := rest.R().Get("http://:8080/domains/ariel17.com.ar")
	assert.Nil(t, err)
	assert.NotNil(t, response)
	assert.Equal(t, http.StatusOK, response.StatusCode())

	// TODO from domain response body to entity representation
	// TODO assert data when domain is found
}
