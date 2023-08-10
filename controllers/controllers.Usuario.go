package controllers

import (
	"SimonBK_Login/db"
	"SimonBK_Login/models"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"gorm.io/gorm"
)

var jwtKey = []byte("tu_clave_secreta")

type LoginInput struct {
	Usuario    string `json:"usuario" binding:"required"`
	Contrasena string `json:"contrasena" binding:"required"`
}

func GenerateAccessToken(usuario *models.Usuario) (string, error) {
	expirationTime := time.Now().Add(24 * time.Hour)
	claims := &jwt.StandardClaims{
		Subject:   usuario.Usuario,
		ExpiresAt: expirationTime.Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func Login(c *gin.Context) {
	var input LoginInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var usuario models.Usuario
	if err := models.GetUsuarioByUsuario(&usuario, input.Usuario); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Usuario o contraseña incorrectos"})
		return
	}

	if err := models.CheckPassword(&usuario, input.Contrasena); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Usuario o contraseña incorrectos"})
		return
	}

	if err := models.CheckPassword(&usuario, input.Contrasena); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Usuario o contraseña incorrectos"})
		return
	}
	// Crear un token JWT
	tokenString, err := GenerateAccessToken(&usuario)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al firmar el token"})
		return
	}
	//Definir refreshToken aqui, antes de las condiciones
	var refreshToken *models.RefreshToken
	// Busca un refreshToken existente para ese usuario.
	var existingToken models.RefreshToken
	err = db.DBConn.Where("usuario_id = ?", usuario.ID).First(&existingToken).Error

	// Si el token existe, actualízalo
	if err == nil {
		// Crear un nuevo token de refresco
		refreshToken, err = models.GenerateRefreshToken(&usuario)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al generar el token de refresco"})
			return
		}

		// Actualiza el token existente
		err = db.DBConn.Model(&existingToken).Updates(models.RefreshToken{Token: refreshToken.Token, ExpiryDate: refreshToken.ExpiryDate}).Error
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al actualizar el token de refresco"})
			return
		}
	} else if err == gorm.ErrRecordNotFound {
		// Si el token no existe, crea uno nuevo
		refreshToken, err := models.GenerateRefreshToken(&usuario)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al generar el token de refresco"})
			return
		}

		// Guarda el nuevo token en la base de datos
		err = db.DBConn.Save(&refreshToken).Error
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al guardar el nuevo token de refresco"})
			return
		}

		// Antes de enviar la respuesta, verifica si refreshToken es nil
		if refreshToken == nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error interno, refreshToken es nulo"})
			return
		}

	} else {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al buscar el token de refresco"})
		return
	}

	// estructura del permiso para responder
	type PermisoResponse struct {
		ID         uint
		FkUsuario  uint
		FkRoles    uint
		FkModulo   uint
		Lectura    bool
		Escritura  bool
		Eliminar   bool
		Actualizar bool
	}

	// obtener los permisos del usuario
	var permisos []PermisoResponse
	db.DBConn.Table("permiso_usuarios").Where("fk_usuario = ?", usuario.ID).Scan(&permisos)

	c.JSON(http.StatusOK, gin.H{
		"accessToken":  tokenString,
		"refreshToken": refreshToken.Token,
		"message":      "Inicio de sesión exitoso",
		"id_usuario":   usuario.ID,
		"nombre":       usuario.Nombre,
		"id_compania":  usuario.FkCompania,
		"id_cliente":   usuario.FkCliente,
		"permisos":     permisos,
	})

}
