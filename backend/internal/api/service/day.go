package service

import (
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/VictorzllDev/ZorpPay/backend/internal/api/repository"
	"github.com/VictorzllDev/ZorpPay/backend/internal/domain/entities"
	"gorm.io/gorm"
)

type DayService interface {
	CreateDay(day *entities.Day) error
}

type dayService struct {
	repository repository.DayRepository
}

func NewDayService(repository repository.DayRepository) DayService {
	return &dayService{repository: repository}
}

func (s *dayService) CreateDay(day *entities.Day) error {
	if day.UserID == 0 {
		return errors.New("user ID é obrigatório")
	}

	if day.Date.IsZero() {
		return errors.New("data inválida")
	}

	if day.Date.After(time.Now()) {
		return errors.New("não é permitido datas futuras")
	}

	// Verifica se o dia já existe para o usuário
	existingDay, err := s.repository.FindByUserAndDate(day.UserID, day.Date)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return errors.New("erro interno ao verificar data")
	}

	if existingDay != nil {
		return fmt.Errorf("data %s já registrada", day.Date.Format("02-01-2006"))
	}

	// Salva o novo dia
	if err := s.repository.Save(day); err != nil {
		log.Printf("Erro ao salvar dia: %v", err)
		return errors.New("erro interno ao registrar data")
	}

	return nil
}
