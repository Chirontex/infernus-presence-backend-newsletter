package validator

import (
	"fmt"
	"regexp"
	"strings"
)

var emailRegex = regexp.MustCompile(`^[a-zA-Z0-9._%+\-]+@[a-zA-Z0-9.\-]+\.[a-zA-Z]{2,}$`)

type ValidationError struct {
	Field   string
	Message string
}

func (e ValidationError) Error() string {
	return fmt.Sprintf("%s: %s", e.Field, e.Message)
}

func ValidateEmail(email string) error {
	email = strings.TrimSpace(email)

	if email == "" {
		return ValidationError{
			Field:   "email",
			Message: "email is required",
		}
	}

	if len(email) > 254 {
		return ValidationError{
			Field:   "email",
			Message: "email is too long (max 254 characters)",
		}
	}

	if !emailRegex.MatchString(email) {
		return ValidationError{
			Field:   "email",
			Message: "invalid email format",
		}
	}

	return nil
}

func ValidateClientToken(token, expectedToken string) error {
	token = strings.TrimSpace(token)

	if token == "" {
		return ValidationError{
			Field:   "clientToken",
			Message: "clientToken is required",
		}
	}

	if token != expectedToken {
		return ValidationError{
			Field:   "clientToken",
			Message: "invalid clientToken",
		}
	}

	return nil
}
