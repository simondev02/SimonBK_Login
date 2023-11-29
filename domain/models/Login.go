package models

import "gorm.io/gorm"

type Login struct {
	gorm.Model
	FkUser    uint   `gorm:"column:fk_user"`
	IpAddress string `gorm:"column:ip_address"`
	UserAgent string `gorm:"column:user_agent"`
	Success   bool   `gorm:"column:success"`
}
