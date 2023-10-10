package swagger

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
