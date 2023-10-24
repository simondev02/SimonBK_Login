package controllers

import (
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
	RoleId     uint `json:"roleId"`
}

func GenerateAccessToken(user *models.User) (string, error) {

	jwtKey := os.Getenv("JWT_KEY")

	expirationTime := time.Now().Add(24 * time.Hour)
	claims := &CustomClaims{
		StandardClaims: jwt.StandardClaims{
			Subject:   user.Username,
			ExpiresAt: expirationTime.Unix(),
		},
		FkCompany:  user.Fk_Company,
		FkCustomer: user.Fk_Customer,
		UserId:     user.ID,
		RoleId:     user.Fk_Role,
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
// @Param login body LoginInput true "Credenciales del usuario"
// @Success 200 {object} swagger.LoginResponse "Respuesta exitosa con tokens y detalles del usuario"
// @Failure 400 {object} map[string]string "Error: Datos inválidos"
// @Failure 401 {object} map[string]string "Error: Usuario o contraseña incorrectos"
// @Failure 500 {object} map[string]string "Error interno del servidor"
// @Router /users/login/ [post]
func Login(c *gin.Context) {
	var input LoginInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var username models.UserDetail
	if err := models.GetUserDetail(&username, input.Username); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Usuario o contraseña incorrectos"})
		return
	}

	if err := models.CheckPassword(&username, input.Password); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Usuario o contraseña incorrectos"})
		return
	}

	var user models.User
	// Crear un token JWT
	tokenString, err := GenerateAccessToken(&user)
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
		refreshToken, err = models.GenerateRefreshToken(&user)
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
		refreshToken, err := models.GenerateRefreshToken(&user)
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

	c.JSON(http.StatusOK, gin.H{
		"accessToken":  tokenString,
		"refreshToken": refreshToken.Token,
		"message":      "Inicio de sesión exitoso",
		"userId":       username.ID,
		"username":     username.Username,
		"name":         username.Name,
		"id_company":   username.Fk_Company,
		"id_customer":  username.Fk_Customer,
		"roleId":       username.Fk_Role,
		"role":         username.RoleDescription,
	})
}
