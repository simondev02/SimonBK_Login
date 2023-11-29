package services

import (
	"SimonBK_Login/domain/models"
	"SimonBK_Login/infra/db"
	"errors"
	"time"

	"github.com/google/uuid"
)

func UpdateRefreshToken(userId uint) (string, error) {
	// Genera un nuevo token
	newToken := uuid.New().String()
	ExpiryDate := time.Now().Add(time.Hour * 24 * 7)
	// Actualiza el token en la base de datos
	result := db.DBConn.Model(&models.RefreshToken{}).Where("fk_user = ?", userId).Update("token", newToken).Update("expiry_date", ExpiryDate)
	// Verifica si la operación de actualización fue exitosa
	if result.Error != nil {
		// Retorna el error si algo salió mal
		return "", result.Error
	}
	if result.RowsAffected == 0 {
		// Ninguna fila fue actualizada, puede que no exista un token para este usuario
		return "", errors.New("ningún token de refresco fue actualizado")
	}
	// Todo salió bien, retorna nil
	return newToken, nil
}
