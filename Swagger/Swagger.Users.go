// swagger/Swagger.User.go
package swagger

// Users represents a users model
// swagger:model

type LoginInput struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type LoginResponse struct {
	AccessToken  string               `json:"accessToken"`
	RefreshToken string               `json:"refreshToken"`
	Message      string               `json:"message"`
	IDUsername   uint                 `json:"id_username"`
	Name         string               `json:"name"`
	IDCompany    int                  `json:"id_company"`
	IDCustomer   int                  `json:"id_customer"`
	Permission   []PermissionResponse `json:"permission"`
}

type PermissionResponse struct {
	ID         uint `json:"id"`
	FkUsername uint `json:"fk_username"`
	FkRole     uint `json:"fk_role"`
	FkModule   uint `json:"fk_module"`
	Read       bool `json:"read"`
	Write      bool `json:"write"`
	Delete     bool `json:"delete"`
	Update     bool `json:"update"`
}
