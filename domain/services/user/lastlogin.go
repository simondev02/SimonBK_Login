package services

import (
	"SimonBK_Login/domain/models"
	"SimonBK_Login/infra/db"
	"time"
)

func UpdateLastLogin(userId uint) (string, error) {
	var user models.UsersDevs

	// Actualizar el campo last_login en la base de datos
	now := time.Now()
	err := db.DBConn.Model(&user).Where("id = ?", userId).Update("last_login", now).Error
	if err != nil {
		return "", err
	}

	// Recargar los datos del usuario para obtener la fecha y hora actualizadas
	err = db.DBConn.Where("id = ?", userId).First(&user).Error
	if err != nil {
		return "", err
	}

	// Formatear la fecha y hora para devolverla
	formattedTime := user.Last_login.Format("2006-01-02 15:04:05")
	return formattedTime, nil
}
