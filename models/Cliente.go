package models

import (
	"gorm.io/gorm"
)

type Cliente struct {
	gorm.Model
	Nombre               string
	Direccion            string
	Telefono             string
	Email                string
	NumeroIdentificacion string
	FkCompania           int
	Compania             Compania  `gorm:"foreignKey:FkCompania"`
	Usuario              []Usuario `gorm:"foreignKey:FkCliente"`
}
