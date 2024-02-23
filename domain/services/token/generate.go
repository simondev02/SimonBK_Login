package services

func GenerateRefreshToken(userId uint) (string, error) {
	var Token string
	_, err := ValidateRefreshTokenByUserId(userId)
	if err != nil {
		if err.Error() == "refresh token no encontrado" {
			Token, err = CreateRefreshToken(userId)
			if err != nil {
				return "", err
			}
		} else {
			// Si hay un error diferente al no encontrar el token
			return "", err
		}
	} else {
		Token, err = UpdateRefreshToken(userId)
		if err != nil {
			return "", err
		}
	}
	return Token, nil
}
