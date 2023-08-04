package models

import (
	"gorm.io/gorm"
)

type PermisoUsuario struct {
	gorm.Model
	FkUsuario  int `gorm:"uniqueIndex:idx_user_module"`
	FkRoles    int
	FkModulo   int `gorm:"uniqueIndex:idx_user_module"`
	Lectura    bool
	Escritura  bool
	Eliminar   bool
	Actualizar bool
	Usuario    Usuario `gorm:"foreignKey:FkUsuario"`
	Roles      Roles   `gorm:"foreignKey:FkRoles"`
	Modulo     Modulo  `gorm:"foreignKey:FkModulo"`
}
