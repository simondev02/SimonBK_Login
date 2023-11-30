package controllers

import (
	"SimonBK_Login/api/views"
	login "SimonBK_Login/domain/services/login"
	services "SimonBK_Login/domain/services/password"
	token "SimonBK_Login/domain/services/token"
	userServices "SimonBK_Login/domain/services/user"
	"SimonBK_Login/utilities"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Login godoc
// @Summary Iniciar sesión
// @Description Autentica a un usuario y devuelve un token de acceso y un token de refresco.
// @Tags autenticación
// @Accept  json
// @Produce  json
// @Param   login  body      views.LoginForm  true  "Credenciales del usuario"
// @Success 200 {object} views.Response "Respuesta exitosa con tokens y detalles del usuario"
// @Failure 400 "Error: Datos inválidos o problema con el formato del email"
// @Failure 401 "Error: Credenciales incorrectas o intento de inicio de sesión fallido"
// @Failure 500 "Error interno del servidor"
// @Router /users/login/ [post]
func Login(c *gin.Context) {
	var input views.LoginForm
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "error": err.Error()})
		return
	}
	// Validar el formato del email
	if err := utilities.CheckEmail(input.Email); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "message": err.Error()})
		return
	}
	clientIP := c.ClientIP()
	userAgent := c.GetHeader("User-Agent")

	user, err := userServices.GetUserByEmail(input.Email)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "message": "Error al obtener detalles del usuario"})
		return
	}

	// Verificar credenciales
	ok, err := services.ValidatePassword(input.Email, input.Password)
	if !ok || err != nil {
		attempt := user.Login_attempts + 1
		err := userServices.IncrementLoginAttempt(user.ID, attempt)
		if err != nil {
			return
		}
		err = login.CreateLogin(user.ID, clientIP, userAgent, false)
		c.JSON(http.StatusUnauthorized, gin.H{"failedattempt": attempt, "success": false, "message": "Credenciales incorrectas"})
		return
	}

	// Crear un token JWT
	tokenString, err := utilities.GenerateAccessToken(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "message": "Error al firmar el token"})
		return
	}
	// Crear un token de actualización
	refreshToken, err := token.GenerateRefreshToken(user.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "message": "Error al generar el token de refresco"})
		return
	}
	// Actualizar el último inicio de sesión
	last_log, err := userServices.UpdateLastLogin(user.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "message": "Error al actualizar el último inicio de sesión"})
		return
	}
	err = login.CreateLogin(user.ID, clientIP, userAgent, true)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "message": "Error al registrar inicio de sesión"})
		return
	}
	user.Last_login = last_log
	// Envía la respuesta
	response := views.Response{
		Success:        true,
		AccessToken:    tokenString,
		FailedAttempts: user.Login_attempts,
		RefreshToken:   refreshToken,
		Message:        "Inicio de sesión exitoso",
		User:           user,
	}
	c.JSON(http.StatusOK, response)
}
