package email

import (
	"strings"
)

func format(validatedEmail string) string {
	return strings.ToLower(validatedEmail)
}
