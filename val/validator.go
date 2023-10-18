package val

import (
	"fmt"
	"net/mail"
	"regexp"
)

var (
	isValidUsername = regexp.MustCompile(`^[a-z0-9_]+$`).MatchString
	isValidFullName = regexp.MustCompile(`^[a-zA-Z\\s]+$`).MatchString
)

func validateString(value string, minLengt int, maxLength int) error {
	n := len(value)
	if n < minLengt || n > maxLength {
		return fmt.Errorf("must contain from %d-%d character", minLengt, maxLength)
	}

	return nil
}

func ValidateUsername(value string) error {
	if err := validateString(value, 3, 100); err != nil {
		return err
	}

	if !isValidUsername(value) {
		return fmt.Errorf("must contain only lowercase letter, digits, or underscore")
	}
	return nil
}

func ValidateFullName(value string) error {
	if err := validateString(value, 3, 100); err != nil {
		return err
	}

	if !isValidFullName(value) {
		return fmt.Errorf("must contain only letter, digits, or spaces")
	}
	return nil
}

func ValidatePassword(value string) error {
	return validateString(value, 6, 100)
}

func ValidateEmail(value string) error {
	if err := validateString(value, 3, 200); err != nil {
		return err
	}

	if _, err := mail.ParseAddress(value); err != nil {
		return fmt.Errorf("is not valid email address")
	}
	return nil
}
