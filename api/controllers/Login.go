package controllers

import (
	"SimonBK_Login/api/views"
	services "SimonBK_Login/domain/services/password"
	token "SimonBK_Login/domain/services/token"
	user "SimonBK_Login/domain/services/user"
	"SimonBK_Login/utilities"
	"net/http"

	"github.com/gin-gonic/gin"
)

// @Summary Iniciar sesión
// @Description Autentica a un usuario y devuelve un token de acceso y un token de refresco
// @Accept json
// @Produce json
// @Param login body LoginInput true "Credenciales del usuario"
// @Success 200 {object} swagger.LoginResponse "Respuesta exitosa con tokens y detalles del usuario"
// @Failure 400 {object} map[string]string "Error: Datos inválidos"
// @Failure 401 {object} map[string]string "Error: Usuario o contraseña incorrectos"
// @Failure 500 {object} map[string]string "Error interno del servidor"
// @Router /users/login/ [post]
func Login(c *gin.Context) {
	var input views.LoginForm
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// Validar el formato del email
	if err := utilities.CheckEmail(input.Email); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"succes": false, "message": err.Error()})
		return
	}
	// Verificar credenciales
	ok, err := services.ValidatePassword(input.Email, input.Password)
	if ok != "Contraseña correcta" || err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"success": false, "message": "Credenciales incorrectas"})

		return
	}
	user, err := user.GetUserInfo(input.Email)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "message": "Error al obtener detalles del usuario"})
		return
	}
	// Crear un token JWT
	tokenString, err := utilities.GenerateAccessToken(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "message": "Error al firmar el token"})
		return
	}
	// Manejar el token de refresco
	var refreshToken string
	err = token.ValidateRefreshToken(user.ID)
	if err != nil {
		if err.Error() == "refresh token no encontrado" {
			// Si el token no existe, crea uno nuevo
			refreshToken, err = token.CreateRefreshToken(user.ID)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"success": false, "message": "Error al generar el token de refresco"})
				return
			}
		} else {
			// Si hay un error diferente al no encontrar el token
			c.JSON(http.StatusInternalServerError, gin.H{"success": false, "message": "Error al validar el token de refresco"})
			return
		}
	} else {
		// Si el token existe, actualízalo
		refreshToken, err = token.UpdateRefreshToken(user.ID)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"success": false, "message": "Error al actualizar el token de refresco"})
			return
		}
	}
	// Envía la respuesta
	response := views.Response{
		Success:      true,
		AccessToken:  tokenString,
		Attempt:      0,
		RefreshToken: refreshToken,
		Message:      "Inicio de sesión exitoso",
		User:         user,
	}
	c.JSON(http.StatusOK, response)
}
