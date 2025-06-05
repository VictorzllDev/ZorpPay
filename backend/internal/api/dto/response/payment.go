package response

type Payment struct {
	ID   int    `json:"id"`
	Date string `json:"date" example:"02-01-2025"`
}
