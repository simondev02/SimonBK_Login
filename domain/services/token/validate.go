package services

import (
	"SimonBK_Login/domain/models"
	"SimonBK_Login/infra/db"
	"errors"

	"gorm.io/gorm"
)

func ValidateRefreshToken(userId uint) error {
	var RefreshToken models.RefreshToken
	result := db.DBConn.Where("fk_user = ?", userId).First(&RefreshToken)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			// Retorna un error específico si el token no se encuentra
			return errors.New("refresh token no encontrado")
		}
		return result.Error
	}

	return nil
}
