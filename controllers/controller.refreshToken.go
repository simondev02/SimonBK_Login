package controllers

import (
	"SimonBK_Login/db"
	"SimonBK_Login/models"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// @Summary Refrescar el token de acceso
// @Description Refresca un token de acceso utilizando un token de refresco válido
// @Accept json
// @Produce json
// @Param refreshToken body swagger.RefreshTokenInput true "Token de refresco"
// @Success 200 {object} swagger.accessTokenResponse "Respuesta exitosa con un nuevo token de acceso"
// @Failure 400 {object} map[string]string "Error: Datos inválidos"
// @Failure 401 {object} map[string]string "Error: Token inválido o expirado"
// @Failure 500 {object} map[string]string "Error interno del servidor"
// @Router /user/refresh [post]
func Refresh(c *gin.Context) {
	var input struct {
		RefreshToken string `json:"refreshToken" binding:"required"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Busca el token de actualización en la base de datos
	var refreshToken models.RefreshToken
	if err := db.DBConn.Where("token = ?", input.RefreshToken).First(&refreshToken).Error; err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Token inválido"})
		return
	}

	// Comprueba la fecha de caducidad
	if refreshToken.ExpiryDate.Before(time.Now()) {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Token expirado"})
		return
	}

	// Obtiene el usuario asociado con el token de actualización
	var user models.User
	if err := db.DBConn.First(&user, refreshToken.FkUser).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al obtener el usuario"})
		return
	}

	// Genera un nuevo token de acceso
	accessToken, err := GenerateAccessToken(&user) // Asume que ya tienes esta función
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"accessToken": accessToken,
	})
}
