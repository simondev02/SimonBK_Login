package models

import (
	"gorm.io/gorm"
)

type UserPermission struct {
	gorm.Model
	FkUser   int `gorm:"uniqueIndex:idx_user_module"`
	FkRole   int
	FkModule int `gorm:"uniqueIndex:idx_user_module"`
	Read     bool
	Write    bool
	Delete   bool
	Update   bool
	User     User   `gorm:"foreignKey:FkUser"`
	Role     Role   `gorm:"foreignKey:FkRole"`
	Module   Module `gorm:"foreignKey:FkModule"`
}
