package models

import (
	"SimonBK_Login/db"
	"errors"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type Usuario struct {
	gorm.Model
	Nombre         string
	Usuario        string
	Contrasena     string
	FkCompania     int
	FkCliente      int
	Compania       Compania         `gorm:"foreignKey:FkCompania"`
	Cliente        Cliente          `gorm:"foreignKey:FkCliente"`
	PermisoUsuario []PermisoUsuario `gorm:"foreignKey:FkUsuario"`
}

// GetUsuarioByUsuario busca un usuario por nombre de usuario en la base de datos
func GetUsuarioByUsuario(usuario *Usuario, nombreUsuario string) (err error) {
	if err = db.DBConn.Where("usuario = ?", nombreUsuario).First(usuario).Error; err != nil {
		return err
	}
	return nil
}
func CheckPassword(usuario *Usuario, contrasena string) error {
	// Comprueba si la contraseña proporcionada coincide con la contraseña almacenada.
	err := bcrypt.CompareHashAndPassword([]byte(usuario.Contrasena), []byte(contrasena))

	if err != nil {
		// Si hay un error, la contraseña no coincide
		return errors.New("contraseña incorrecta")
	}

	// No hubo error, así que las contraseñas coinciden
	return nil
}
