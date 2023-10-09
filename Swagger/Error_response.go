// Swagger/error_response.go
package swagger

// ErrorResponse representa una respuesta de error estándar.
// swagger:model ErrorResponse
type ErrorResponse struct {
	Message string `json:"message" example:"Error message"`
	Code    int    `json:"code" example:"500"`
}
