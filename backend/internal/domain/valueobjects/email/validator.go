package email

import (
	"regexp"
	"strings"
)

var (
	emailRegex = regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
)

func validate(rawEmail string) (string, error) {
	trimmed := strings.TrimSpace(rawEmail)
	if !emailRegex.MatchString(trimmed) {
		return "", ErrInvalidEmail
	}
	return trimmed, nil
}
