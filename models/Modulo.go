package models

import (
	"gorm.io/gorm"
)

type Modulo struct {
	gorm.Model
	Nombre      string
	Descripcion string
}
