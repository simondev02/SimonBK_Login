package views

type User struct {
	ID         uint
	FkCompany  int
	FkCustomer int
	Email      string
	Name       string
	Lastname   string
	FkRole     int
	Role       string
}
