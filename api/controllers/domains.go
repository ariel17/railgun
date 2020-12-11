package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/ariel17/railgun/api/services"
)

// GetDomainController handles the read endpoint request to show a domain data
// if exists.
func GetDomainController(c *gin.Context) {
	value := c.Param("value")
	domain, err := services.GetDomain(value)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	if domain == nil {
		c.AbortWithStatus(http.StatusNotFound)
		return
	}
	c.JSON(http.StatusOK, domain)
}
