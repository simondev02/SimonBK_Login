package services

import (
	"SimonBK_Login/domain/models"
	"SimonBK_Login/infra/db"
	"errors"
	"log"

	"gorm.io/gorm"
)

func ValidateRefreshTokenByUserId(userId uint) (uint, error) {
	var RefreshToken models.RefreshToken

	log.Println("Obteniendo refresh token por fk_user")
	Query := db.DBConn.Where("fk_user = ?", userId).First(&RefreshToken)
	if Query.Error != nil {
		if errors.Is(Query.Error, gorm.ErrRecordNotFound) {
			return 0, errors.New("refresh token no encontrado")
		}
		return 0, Query.Error
	}

	/* 	curentTime := time.Now()
	   	if RefreshToken.ExpiryDate.Before(curentTime) {
	   		return 0, errors.New("token expirado")
	   	} */
	return RefreshToken.FkUser, nil
}

func GetUserIdByRefreshToken(token string) (uint, error) {
	var RefreshToken models.RefreshToken

	log.Println("Obteniendo refresh token")
	Query := db.DBConn.Where("token = ?", token).First(&RefreshToken)
	if Query.Error != nil {
		if errors.Is(Query.Error, gorm.ErrRecordNotFound) {
			return 0, errors.New("refresh token no encontrado")
		}
		return 0, Query.Error
	}

	return RefreshToken.FkUser, nil
}
