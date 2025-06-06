package service

import (
	"errors"
	"fmt"
	"time"

	"github.com/VictorzllDev/ZorpPay/backend/internal/api/repository"
	"github.com/VictorzllDev/ZorpPay/backend/internal/domain/entities"
)

type PaymentService interface {
	CreatePayment(payment *entities.Payment) error
}

type paymentService struct {
	repository repository.PaymentRepository
}

func NewPaymentService(repository repository.PaymentRepository) PaymentService {
	return &paymentService{repository: repository}
}

func (s *paymentService) CreatePayment(payment *entities.Payment) error {
	if payment.Amount == 0 {
		return errors.New("valor obrigatório")
	}

	if payment.PaymentDate.IsZero() {
		return errors.New("data inválida")
	}

	if payment.PaymentDate.After(time.Now()) {
		return errors.New("não é permitido pagamentos futuros")
	}

	totalDays, err := s.repository.CountUserDays(payment.UserID)
	if err != nil {
		return fmt.Errorf("erro ao buscar dias letivos: %v", err)
	}

	totalUsedInPayments, err := s.repository.GetUserTotalPayments(payment.UserID)
	if err != nil {
		return fmt.Errorf("erro ao buscar pagamentos existentes: %v", err)
	}

	availableBalance := (totalDays * 20) - totalUsedInPayments

	if payment.Amount > availableBalance {
		return fmt.Errorf(
			"saldo insuficiente: R$%d solicitados (equivalente a %d dias), R$%d disponíveis (equivalente a %d dias)",
			payment.Amount,
			payment.Amount/20,
			availableBalance,
			availableBalance/20,
		)
	}

	return s.repository.Save(payment)
}
