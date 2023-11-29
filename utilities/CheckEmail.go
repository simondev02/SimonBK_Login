package utilities

import (
	"errors"
	"regexp"
)

func CheckEmail(email string) error {
	// implementamos Regex para verificar el formato del EMAIL
	emailRegex := regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`)
	if emailRegex.MatchString(email) {
		return nil
	}
	return errors.New("El email no tiene un formato válido")
}
