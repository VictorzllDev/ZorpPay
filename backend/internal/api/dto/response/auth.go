package response

// @Description Authentication response with JWT token
type Auth struct {
	Token string `json:"token"`
}
