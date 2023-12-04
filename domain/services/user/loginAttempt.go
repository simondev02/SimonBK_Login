package services

import (
	"SimonBK_Login/domain/models"
	"SimonBK_Login/infra/db"
)

func IncrementLoginAttempt(userId uint, attempt int) error {
	var user models.Users
	// Realiza la actualización y verifica si hay un error
	result := db.DBConn.Model(&user).Where("id = ?", userId).Update("login_attempts", attempt)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
