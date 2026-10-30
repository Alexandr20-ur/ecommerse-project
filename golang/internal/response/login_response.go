package response

type LoginResponse struct {
	Token string `json:"token"`
	Err   string
	Msg   string
}
