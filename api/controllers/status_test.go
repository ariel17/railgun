package controllers

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/go-resty/resty/v2"
	"github.com/stretchr/testify/assert"
)

func TestPingController(t *testing.T) {
	r := gin.Default()
	r.GET("/ping", PingController)
	ts := httptest.NewServer(r)
	defer ts.Close()

	rest := resty.New()
	response, err := rest.R().Get("http://:8080/ping")
	assert.Nil(t, err)
	assert.NotNil(t, response)
	assert.Equal(t, http.StatusOK, response.StatusCode())
}
