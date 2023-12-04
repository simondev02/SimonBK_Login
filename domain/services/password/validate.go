package services

import (
	"SimonBK_Login/domain/models"
	userServices "SimonBK_Login/domain/services/user"
	"SimonBK_Login/infra/db"
	"SimonBK_Login/utilities"
	"errors"
)

func ValidatePassword(email string, password string) (bool, error) {
	var user models.Users
	// Buscar el usuario por username(email)
	db.DBConn.Where("email = ?", email).First(&user)
	if user.ID == 0 {
		return false, errors.New("Usuario no encontrado")
	}
	// Validar la contraseña
	err := utilities.CheckPasswordHash(password, user.Password)
	if err != nil {
		return false, errors.New("Credenciales incorrectas")
	}
	err = userServices.IncrementLoginAttempt(user.ID, 0)
	if err != nil {
		return false, err
	}
	return true, nil
}
