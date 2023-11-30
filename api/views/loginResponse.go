package views

type Response struct {
	Success        bool   `json:"success"`
	FailedAttempts int    `json:"failedattempts"`
	AccessToken    string `json:"accesstoken"`
	RefreshToken   string `json:"refreshtoken"`
	Message        string `json:"message"`
	User           User   `json:"user"`
}
