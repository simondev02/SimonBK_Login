package services

import (
	"SimonBK_Login/api/views"
	"SimonBK_Login/domain/models"
	"SimonBK_Login/infra/db"
)

func GetUserByEmail(email string) (views.User, error) {
	var User models.Users

	err := db.DBConn.Preload("Contacts").Preload("Role").Where("email = ?", email).First(&User).Error
	if err != nil {
		return views.User{}, err
	}
	user := views.User{
		ID:             User.ID,
		FkCompany:      User.FkCompany,
		FkCustomer:     User.FkCustomer,
		Email:          User.Email,
		Name:           User.Contacts[0].Name,
		Lastname:       User.Contacts[0].Lastname,
		FkRole:         User.FkRole,
		Role:           User.Role.RoleDescription,
		Last_login:     User.Last_login.String(),
		Login_attempts: User.Login_attempts,
	}
	return user, nil
}

func GetUserById(id uint) (views.User, error) {
	var User models.Users

	err := db.DBConn.Preload("Contacts").Preload("Role").Where("id = ?", id).First(&User).Error
	if err != nil {
		return views.User{}, err
	}
	user := views.User{
		ID:         User.ID,
		FkCompany:  User.FkCompany,
		FkCustomer: User.FkCustomer,
		Email:      User.Email,
		Name:       User.Contacts[0].Name,
		Lastname:   User.Contacts[0].Lastname,
		FkRole:     User.FkRole,
		Role:       User.Role.RoleDescription,
	}
	return user, nil
}
