package models

import (
	"gorm.io/gorm"
)

type Role struct {
	gorm.Model
	Name            string `gorm:"uniqueIndex:idx_name_fkcompany"`
	Description     string
	FkCompany       uint    `gorm:"uniqueIndex:idx_name_fkcompany"`
	Company         Company `gorm:"foreignKey:FkCompany"`
	DeletedByUserID *uint
	UpdatedByUserID *uint
}
