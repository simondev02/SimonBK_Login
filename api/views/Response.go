package views

type Response struct {
	Success      bool
	Attempt      int
	AccessToken  string
	RefreshToken string
	Message      string
	User         User
}
