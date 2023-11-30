package models

import "gorm.io/gorm"

type Permission struct {
	gorm.Model
	Permission string `gorm:"column:permission"`
}
