package models

import (
	"gorm.io/gorm"
)

type Company struct {
	gorm.Model
	Name                 string `gorm:"unique"`
	Address              string
	Phone                string
	Email                string
	IdentificationNumber string `gorm:"unique"`
	DeletedByUserID      *uint
	UpdatedByUserID      *uint
}
