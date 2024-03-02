package services

import (
	"SimonBK_Login/api/views"
	"SimonBK_Login/domain/models"
	"SimonBK_Login/infra/db"
)

func GetUserById(userId uint) (views.User, error) {
	var user models.Users

	// Carga el usuario
	err := db.DBConn.Where("id = ?", userId).First(&user).Error
	if err != nil {
		return views.User{}, err
	}

	// Carga el UserRole asociado al usuario
	var userRole models.UserRole
	err = db.DBConn.Where("fk_user = ?", user.ID).First(&userRole).Error
	if err != nil {
		return views.User{}, err
	}

	// Carga el CompanyRole usando FkRole de UserRole
	var companyRole models.CompanyRole
	err = db.DBConn.Where("id = ?", userRole.FkRole).First(&companyRole).Error
	if err != nil {
		return views.User{}, err
	}

	// Carga el Role asociado con CompanyRole
	var role models.Role
	err = db.DBConn.Where("id = ?", companyRole.FkRole).First(&role).Error
	if err != nil {
		return views.User{}, err
	}

	// // Carga los contactos del usuario
	// err = db.DBConn.Model(&user).Association("Contacts").Find(&user.Contacts)
	// if err != nil {
	// 	return views.User{}, err
	// }

	// Construye el objeto views.User
	userView := views.User{
		ID:             user.ID,
		FkCompany:      user.FkCompany,
		FkCustomer:     user.FkCustomer,
		Email:          user.Email,
		Name:           "",
		Lastname:       "",
		FkRole:         role.ID,
		Role:           role.RoleDescription,
		Last_login:     user.Last_login.String(),
		Login_attempts: user.Login_attempts,
	}

	// // Asignar nombre y apellido de los contactos, si existen
	// if len(user.Contacts) > 0 {
	// 	userView.Name = user.Contacts[0].Name
	// 	userView.Lastname = user.Contacts[0].Lastname
	// }

	return userView, nil
}
