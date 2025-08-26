package request

// @Description Authentication request
type SignIn struct {
	Email    string `json:"email" example:"johndoe@example.com" binding:"required,email,max=50"`
	Password string `json:"password" example:"12345678" binding:"required,min=8,max=30"`
}

// @Description Authentication request
type SignUp struct {
	Name     string `json:"name" example:"John Doe" binding:"required,min=3,max=30"`
	Email    string `json:"email" example:"johndoe@example.com" binding:"required,email,max=50"`
	Password string `json:"password" example:"12345678" binding:"required,min=8,max=30"`
}
