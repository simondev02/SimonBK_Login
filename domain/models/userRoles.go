package models

import "gorm.io/gorm"

type UserRole struct {
	gorm.Model
	FkUser      int         `gorm:"column:fk_user"`
	FkRole      int         `gorm:"column:fk_role"` // Relación hacia tabla company_roles
	CompanyRole CompanyRole `gorm:"foreignKey:FkRole"`
}
