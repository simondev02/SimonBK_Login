package views

type User struct {
	ID             uint   `json:"id"`
	FkCompany      int    `json:"fkcompany"`
	FkCustomer     int    `json:"fkcustomer"`
	Email          string `json:"email"`
	Name           string `json:"name"`
	Lastname       string `json:"lastname"`
	FkRole         int    `json:"fkrole"`
	Role           string `json:"role"`
	Login_attempts int    `json:"-"`
	Last_login     string `json:"lastlogin"`
}
