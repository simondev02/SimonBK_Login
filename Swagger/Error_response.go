package swagger

type ErrorResponse struct {
	Message string `json:"message" example:"Error message"`
	Code    int    `json:"code" example:"500"`
}
