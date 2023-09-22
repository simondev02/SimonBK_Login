package models

import (
	"gorm.io/gorm"
)

type UserPermission struct {
	gorm.Model
	FkRole          int
	FkModule        int
	Read            bool
	Write           bool
	Delete          bool
	Update          bool
	DeletedByUserID *uint
	UpdatedByUserID *uint
	Role            Role   `gorm:"foreignKey:FkRole"`
	Module          Module `gorm:"foreignKey:FkModule"`
}
