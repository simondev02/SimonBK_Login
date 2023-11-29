package models

import "gorm.io/gorm"

type UserContact struct {
	gorm.Model
	FkUserDev *uint `gorm:"column:fk_user_dev"`
	FkContact uint  `gorm:"column:fk_contact"`
}
