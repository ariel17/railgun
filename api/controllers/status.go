package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// PingController serves an status code 200 as proof of service availability.
// @Summary HTTP 200 OK header as availavility proof.
// @Description Returns an HTTP 200 OK header.
// @Router /ping [get]
func PingController(c *gin.Context) {
	c.Writer.WriteHeader(http.StatusOK)
}
