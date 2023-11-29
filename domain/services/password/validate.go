package services

import (
	"SimonBK_Login/domain/models"
	"SimonBK_Login/infra/db"
	"SimonBK_Login/utilities"
	"errors"
)

func ValidatePassword(email string, password string) (string, error) {
	var user models.UsersDevs

	// Buscar el usuario por username(email)
	db.DBConn.Where("email = ?", email).First(&user)
	if user.ID == 0 {
		return "", errors.New("Usuario no encontrado")
	}

	// Validar la contraseña
	err := utilities.CheckPasswordHash(password, user.Password)
	if err != nil {
		return "", errors.New("Credenciales incorrectas")
	}

	return "Contraseña correcta", nil
}
