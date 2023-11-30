package services

import (
	"SimonBK_Login/domain/models"
	"SimonBK_Login/infra/db"
	"errors"
	"time"

	"gorm.io/gorm"
)

func ValidateRefreshToken(param interface{}) (uint, error) {
	var RefreshToken models.RefreshToken
	var Query *gorm.DB

	switch v := param.(type) {
	case string:
		Query = db.DBConn.Where("token = ?", v).First(&RefreshToken)
	case uint:
		Query = db.DBConn.Where("fk_user = ?", v).First(&RefreshToken)
	default:
		return 0, errors.New("tipo de dato no válido")
	}
	if Query.Error != nil {
		if errors.Is(Query.Error, gorm.ErrRecordNotFound) {
			return 0, errors.New("refresh token no encontrado")
		}
		return 0, Query.Error
	}

	curentTime := time.Now()
	if RefreshToken.ExpiryDate.Before(curentTime) {
		return 0, errors.New("token expirado")
	}
	return RefreshToken.FkUser, nil
}
