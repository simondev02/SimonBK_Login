package swagger

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
