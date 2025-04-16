package email

import (
	"errors"
	"fmt"
)

var (
	ErrInvalidEmail = errors.New("invalid email format")
)

type Email struct {
	value string
}

func New(rawEmail string) (Email, error) {
	value, err := validate(rawEmail)
	if err != nil {
		return Email{}, fmt.Errorf("email validation failed: %w", err)
	}
	return Email{value: format(value)}, nil
}

func (e Email) String() string {
	return e.value
}
