package views

type ResourceResponse struct {
	Resource string      `json:"resource"`
	Actions  interface{} `json:"actions"`
}
