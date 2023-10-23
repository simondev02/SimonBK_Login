package models

import (
	"gorm.io/gorm"
)

type Roles struct {
	gorm.Model
	id              int
	roleDescription string
}
