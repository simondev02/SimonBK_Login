package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type RefreshToken struct {
	gorm.Model
	Token           string
	FkUser          uint
	ExpiryDate      time.Time
	DeletedByUserID *uint
	UpdatedByUserID *uint
	User            User `gorm:"foreignKey:FkUser"`
}

func GenerateRefreshToken(user *User) (*RefreshToken, error) {
	// Genera un token de actualización aleatorio
	// (Puedes usar una biblioteca como github.com/google/uuid para esto)
	token := uuid.New().String()

	refreshToken := &RefreshToken{
		Token:      token,
		FkUser:     user.ID,
		ExpiryDate: time.Now().Add(time.Hour * 24 * 7), // Por ejemplo, 7 días
	}

	return refreshToken, nil
}
