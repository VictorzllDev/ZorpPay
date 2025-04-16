package request

type CreateUser struct {
	Name     string `json:"name" binding:"required,min=3,max=30"`
	Email    string `json:"email" binding:"required,email,max=50"`
	Password string `json:"password" binding:"required,min=8,max=30"`
}
