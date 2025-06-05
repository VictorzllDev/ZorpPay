package request

type CreateDay struct {
	Date   string `json:"date" example:"02-01-2025" binding:"required"`
	UserID int    `json:"user_id" binding:"required"`
}
