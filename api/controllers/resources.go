package controllers

import (
	services "SimonBK_Login/domain/services/resources.go"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Resources godoc
// @Summary Get resources by role ID
// @Description Retrieve resources associated with a specific role ID
// @Tags resources
// @Accept  json
// @Produce  json
// @Success 200 {object} views.ResourceResponse "Successfully retrieved resources"
// @Failure 500 {object} map[string]string "Error fetching resources"
// @Router /users/resources [get]
func Resources(c *gin.Context) {
	resources, err := services.GetResourcesMap()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, resources)
}
