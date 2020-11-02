package controllers

import (
	"context"
	"net/http"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
    "github.com/go-resty/resty/v2"
	"github.com/stretchr/testify/assert"
)

func TestPingController(t *testing.T) {
	r := gin.Default()
	r.GET("/ping", PingController)
	srv := &http.Server{
		Addr: ":8080",
		Handler: r,
	}
	go srv.ListenAndServe()
	c, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer srv.Shutdown(c)
	defer cancel()

	rest := resty.New()
	response, err := rest.R().Get("http://localhost:8080/ping")
	assert.Nil(t, err)
	assert.NotNil(t, response)
	assert.Equal(t, http.StatusOK, response.StatusCode())
}
