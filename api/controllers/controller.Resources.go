package controllers

// Resources godoc
// @Summary Get resources by role ID
// @Description Retrieve resources associated with a specific role ID
// @Tags resources
// @Accept  json
// @Produce  json
// @Success 200 {object} []models.ResourceResponse "Successfully retrieved resources"
// @Failure 500 {object} map[string]string "Error fetching resources"
// @Router /users/resources [get]
// func Resources(c *gin.Context) {
// 	resources, err := models.GetResourcesByRole()
// 	if err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
// 		return
// 	}
// 	c.JSON(http.StatusOK, resources)
// }
