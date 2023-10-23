package controllers

import (
	"SimonBK_Login/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// Resources godoc
// @Summary Get resources by role ID
// @Description Retrieve resources associated with a specific role ID
// @Tags resources
// @Accept  json
// @Produce  json
// @Param roleid path int true "Role ID"
// @Success 200 {object} []models.ResourceResponse "Successfully retrieved resources"
// @Failure 400 {object} map[string]string "Invalid role ID"
// @Failure 500 {object} map[string]string "Error fetching resources"
// @Router /users/resources/{roleid} [get]
func Resources(c *gin.Context) {
	fk_role_str := c.Param("roleid")
	fk_role, err := strconv.Atoi(fk_role_str)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid role ID"})
		return
	}

	resources, err := models.GetResourcesByRole(fk_role)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, resources)
}
