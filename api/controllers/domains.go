package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/ariel17/railgun/api/controllers/presenters"
	"github.com/ariel17/railgun/api/services"
)

// GetDomainController handles the read endpoint request to show a domain data
// if exists.
func GetDomainController(c *gin.Context) {
	value := c.Param("value")
	domain, err := services.GetDomain(value)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if domain == nil {
		c.AbortWithStatus(http.StatusNotFound)
		return
	}
	c.JSON(http.StatusOK, domain)
}

// NewDomainController handles the creation of a new domain entry to be
// verified.
func NewDomainController(c *gin.Context) {
	var newDomain presenters.NewDomain
	if err := c.ShouldBindJSON(&newDomain); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	domain := newDomain.ToDomain()
	err := services.NewDomain(domain)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, domain)
}
