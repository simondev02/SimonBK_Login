package models

import (
	"time"

	"gorm.io/gorm"
)

type RefreshToken struct {
	gorm.Model
	Token           string
	FkUser          uint
	ExpiryDate      time.Time
	DeletedByUserID *uint
	UpdatedByUserID *uint
	User            UsersDevs `gorm:"foreignKey:FkUser"`
}
