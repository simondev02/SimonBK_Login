package models

import "gorm.io/gorm"

type Role struct {
	gorm.Model
	RoleDescription string `gorm:"column:role_description"`
	DeletedByUserID *uint
	UpdatedByUserID *uint
}

type ActionRoles struct {
	Roles []string `json:"roles"`
}
