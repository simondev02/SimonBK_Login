package controllers

import (
	"SimonBK_Login/api/views"
	services "SimonBK_Login/domain/services/token"
	userServices "SimonBK_Login/domain/services/user"
	"SimonBK_Login/utilities"
	"net/http"

	"github.com/gin-gonic/gin"
)

// "SimonBK_Login/db"
// "SimonBK_Login/models"

// @Summary Refrescar el token de acceso
// @Description Refresca un token de acceso utilizando un token de refresco válido
// @Accept json
// @Produce json
// @Param refreshToken body views.RefreshTokenForm true "Token de refresco"
// @Success 200 {object} views.RefreshTokenResponse "Respuesta exitosa con un nuevo token de acceso"
// @Failure 400 {object} map[string]string "Error: Datos inválidos"
// @Failure 401 {object} map[string]string "Error: Token inválido o expirado"
// @Failure 500 {object} map[string]string "Error interno del servidor"
// @Router /auth/refresh [post]
func RefreshTokenHandler(c *gin.Context) {
	// Manejar el token de refresco
	var input views.RefreshTokenForm
	var user views.User

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	// Verificar el token de refresco
	userId, err := services.GetUserIdByRefreshToken(input.RefreshToken)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "message": "Token inválido o expirado"})
		return
	}
	newToken, err := services.UpdateRefreshToken(userId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"sucees": false, "message": "Error al actualizar token de refresco"})
	}

	// Encontrar el usuario asociado con el token de refresco
	u, err := userServices.GetUserById(userId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "message": "Error al obtener detalles del usuario"})
		return
	}
	user = views.User{
		ID:         u.ID,
		FkCompany:  u.FkCompany,
		FkCustomer: u.FkCustomer,
		FkRole:     u.FkRole,
	}

	// Crear un nuevo token de acceso
	tokenString, err := utilities.GenerateAccessToken(user)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "message": "Error al firmar el token"})
		return
	}
	// Enviar la respuesta
	response := views.RefreshTokenResponse{
		AccessToken:  tokenString,
		RefreshToken: newToken,
	}
	c.JSON(http.StatusOK, response)
}
