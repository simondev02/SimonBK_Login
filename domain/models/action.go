package models

import "gorm.io/gorm"

type Action struct {
	gorm.Model
	Action string `gorm:"column:action"`
}
