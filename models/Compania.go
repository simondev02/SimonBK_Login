package models

import (
	"gorm.io/gorm"
)

type Compania struct {
	gorm.Model
	Nombre    string
	Direccion string
	Telefono  string
	Email     string
	Usuario   []Usuario `gorm:"foreignKey:FkCompania"`
	Cliente   []Cliente `gorm:"foreignKey:FkCompania"`
}
