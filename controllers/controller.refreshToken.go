package controllers

import (
	"SimonBK_Login/db"
	"SimonBK_Login/models"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

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
	var usuario models.Usuario
	if err := db.DBConn.First(&usuario, refreshToken.UsuarioID).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al obtener el usuario"})
		return
	}

	// Genera un nuevo token de acceso
	accessToken, err := GenerateAccessToken(&usuario) // Asume que ya tienes esta función
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"accessToken": accessToken,
	})
}
