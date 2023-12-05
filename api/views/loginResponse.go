package views

type Response struct {
	Success        bool   `json:"success"`
	FailedAttempts int    `json:"failedAttempts"`
	AccessToken    string `json:"accessToken"`
	RefreshToken   string `json:"refreshToken"`
	Message        string `json:"message"`
	User           User   `json:"user"`
}
