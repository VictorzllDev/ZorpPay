package response

// @Description Standard error response format for the API
type Error struct {
	Error string `json:"error" example:"Error description"`
}
