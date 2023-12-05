package views

type User struct {
	ID             uint   `json:"id"`
	FkCompany      int    `json:"fkCompany"`
	FkCustomer     int    `json:"fkCustomer"`
	Email          string `json:"email"`
	Name           string `json:"name"`
	Lastname       string `json:"lastname"`
	FkRole         uint   `json:"fkRole"`
	Role           string `json:"role"`
	Login_attempts int    `json:"-"`
	Last_login     string `json:"lastLogin"`
}
