package models

import (
	"SimonBK_Login/db"
	"errors"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name       string
	Username   string
	Password   string
	FkCompany  int
	FkCustomer int
	Company    Company  `gorm:"foreignKey:FkCompany"`
	Customer   Customer `gorm:"foreignKey:FkCustomer"`
}

// GetUsuarioByUsuario busca un usuario por nombre de usuario en la base de datos
func GetUsuarioByUsuario(user *User, nameUser string) (err error) {
	if err = db.DBConn.Where("username = ?", nameUser).First(user).Error; err != nil {
		return err
	}
	return nil
}
func CheckPassword(user *User, password string) error {
	// Comprueba si la contraseña proporcionada coincide con la contraseña almacenada.
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))

	if err != nil {
		// Si hay un error, la contraseña no coincide
		return errors.New("contraseña incorrecta")
	}

	// No hubo error, así que las contraseñas coinciden
	return nil
}
