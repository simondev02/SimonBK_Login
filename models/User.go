package models

import (
	"SimonBK_Login/db"
	"errors"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID              uint
	Username        string
	Password        string
	Name            string
	Fk_Role         uint
	RoleDescription string
	FkCompany       int `gorm:"column:Fk_Company"`
	FkCustomer      int `gorm:"column:Fk_Customer"`
	// Company         Company  `gorm:"foreignKey:FkCompany"`
	// Customer        Customer `gorm:"foreignKey:FkCustomer"`
	// Role            Role     `gorm:"foreignKey:FkRole"`
}

type UserDetail struct {
	ID              uint
	Username        string
	Password        string
	Name            string
	Fk_Role         uint
	RoleDescription string
	Fk_Company      int
	Fk_Customer     int
}

// GetUsuarioByUsuario busca un usuario por nombre de usuario en la base de datos
func GetUserDetail(userDetail *UserDetail, username string) (err error) {
	err = db.DBConn.
		Table("user_contacts").
		Select("users.id, users.username, users.password, contacts.Name || ' ' || contacts.Lastname as name, users.\"Fk_Role\", roles.\"RoleDescription\",  users.\"Fk_Company\", users.\"Fk_Customer\"").
		Joins("INNER JOIN users ON users.id = user_contacts.Fk_User").
		Joins("INNER JOIN contacts ON user_contacts.Fk_Contact = contacts.id").
		Joins("INNER JOIN roles ON users.\"Fk_Role\" = roles.id").
		Where("users.username = ?", username).
		First(userDetail).Error
	return err
}

func GetUserDetailByToken(userDetail *UserDetail, token string) (err error) {
	err = db.DBConn.
		Table("user_contacts").
		Select("users.id, users.username, users.password, contacts.Name || ' ' || contacts.Lastname as name, users.\"Fk_Role\", roles.\"RoleDescription\",  users.\"Fk_Company\", users.\"Fk_Customer\"").
		Joins("INNER JOIN users ON users.id = user_contacts.Fk_User").
		Joins("INNER JOIN contacts ON user_contacts.Fk_Contact = contacts.id").
		Joins("INNER JOIN roles ON users.\"Fk_Role\" = roles.id").
		Joins("INNER JOIN refresh_tokens ON refresh_tokens.fk_User = users.id").
		Where("refresh_tokens.token = ?", token).
		First(userDetail).Error
	return err

}
func CheckPassword(user *UserDetail, password string) error {
	// Comprueba si la contraseña proporcionada coincide con la contraseña almacenada.
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))

	if err != nil {
		// Si hay un error, la contraseña no coincide
		return errors.New("contraseña incorrecta")
	}
	// No hubo error, así que las contraseñas coinciden
	return nil
}
