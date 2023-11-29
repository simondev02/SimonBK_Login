package models

import "gorm.io/gorm"

type Role struct {
	gorm.Model
	RoleDescription string `gorm:"column:RoleDescription"`
	DeletedByUserID *uint
	UpdatedByUserID *uint
}
