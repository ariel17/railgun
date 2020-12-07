package controllers

import (
	"net/http"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/go-resty/resty/v2"
	"github.com/stretchr/testify/assert"

	"github.com/ariel17/railgun/api/tests"
)

func TestPingController(t *testing.T) {
	r := gin.Default()
	r.GET("/ping", PingController)
	server := tests.CreateAndStartTestServer(r)
	defer tests.StopTestServer(server)

	rest := resty.New()
	response, err := rest.R().Get("http://:8080/ping")
	assert.Nil(t, err)
	assert.NotNil(t, response)
	assert.Equal(t, http.StatusOK, response.StatusCode())
}
