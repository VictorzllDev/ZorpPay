package request

type CreatePayment struct {
	Amount      int    `json:"amount" example:"140" binding:"required"`              // Valor pago (ex: 140)
	PaymentDate string `json:"payment_date" example:"02-01-2025" binding:"required"` // Data do pagamento
	UserID      int    `json:"user_id" binding:"required"`                           // Chave estrangeira
}
