package controllers

import (
	swagger "SimonBK_Login/Swagger"
	"SimonBK_Login/db"
	"SimonBK_Login/models"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"gorm.io/gorm"
)

type LoginInput struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}
type CustomClaims struct {
	jwt.StandardClaims
	FkCompany  int  `json:"fk_company"`
	FkCustomer int  `json:"fk_customer"`
	UserId     uint `json:"userId"`
}

func GenerateAccessToken(user *models.User) (string, error) {

	jwtKey := os.Getenv("JWT_KEY")

	expirationTime := time.Now().Add(24 * time.Hour)
	claims := &CustomClaims{
		StandardClaims: jwt.StandardClaims{
			Subject:   user.Username,
			ExpiresAt: expirationTime.Unix(),
		},
		FkCompany:  user.FkCompany,
		FkCustomer: user.FkCustomer,
		UserId:     user.ID,
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(jwtKey))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

// @Summary Iniciar sesión
// @Description Autentica a un usuario y devuelve un token de acceso y un token de refresco
// @Accept json
// @Produce json
// @Param login body swagger.LogingInput true "Credenciales del usuario"
// @Success 200 {object} swagger.LoginResponse "Respuesta exitosa con tokens y detalles del usuario"
// @Failure 400 {object} map[string]string "Error: Datos inválidos"
// @Failure 401 {object} map[string]string "Error: Usuario o contraseña incorrectos"
// @Failure 500 {object} map[string]string "Error interno del servidor"
// @Router /user/login/ [post]
func Login(c *gin.Context) {
	var input LoginInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var username models.User
	if err := models.GetUsuarioByUsuario(&username, input.Username); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Usuario o contraseña incorrectos"})
		return
	}

	if err := models.CheckPassword(&username, input.Password); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Usuario o contraseña incorrectos"})
		return
	}
	// Crear un token JWT
	tokenString, err := GenerateAccessToken(&username)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al firmar el token"})
		return
	}
	//Definir refreshToken aqui, antes de las condiciones
	var refreshToken *models.RefreshToken
	// Busca un refreshToken existente para ese usuario.
	var existingToken models.RefreshToken
	err = db.DBConn.Where("fk_user = ?", username.ID).First(&existingToken).Error

	// Si el token existe, actualízalo
	if err == nil {
		// Crear un nuevo token de refresco
		refreshToken, err = models.GenerateRefreshToken(&username)
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
		refreshToken, err := models.GenerateRefreshToken(&username)
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

	// obtener los permisos del usuario
	var permission []swagger.PermissionResponse
	db.DBConn.Table("user_permissions").Where("fk_user = ?", username.ID).Scan(&permission)

	c.JSON(http.StatusOK, gin.H{
		"accessToken":  tokenString,
		"refreshToken": refreshToken.Token,
		"message":      "Inicio de sesión exitoso",
		"id_username":  username.ID,
		"name":         username.Name,
		"id_company":   username.FkCompany,
		"id_customer":  username.FkCustomer,
		"permission":   permission,
	})

}
