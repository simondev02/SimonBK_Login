package models

import (
	"gorm.io/gorm"
)

type Roles struct {
	gorm.Model
	Nombre      string
	Descripcion string
}
