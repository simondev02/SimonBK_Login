package services

import (
	"SimonBK_Login/domain/models"
	"SimonBK_Login/infra/db"
	"time"

	"github.com/google/uuid"
)

func CreateRefreshToken(userId uint) (string, error) {

	refreshToken := models.RefreshToken{
		FkUser:     userId,
		Token:      uuid.New().String(),
		ExpiryDate: time.Now().Add(time.Hour * 24 * 7),
	}

	err := db.DBConn.Create(&refreshToken).Error
	if err != nil {
		return "", err
	}

	return refreshToken.Token, nil
}
