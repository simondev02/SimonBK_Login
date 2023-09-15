package models

import (
	"gorm.io/gorm"
)

type Module struct {
	gorm.Model
	Name            string
	Description     string
	DeletedByUserID *uint
	UpdatedByUserID *uint
}
