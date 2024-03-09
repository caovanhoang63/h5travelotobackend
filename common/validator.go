package common

import (
	"github.com/asaskevich/govalidator"
	"regexp"
	"strings"
)

func IsEmail(email string) bool {
	return govalidator.IsEmail(email)
}

func MinLength(str string, minlength int) bool {
	return len(str) >= minlength
}

func IsEmpty(str string) bool {
	return strings.TrimSpace(str) == ""
}

// IsValidPassword checks if the password is valid,
// Minimum eight characters, at least one letter, one number and one special character

func IsValidPassword(password string) bool {
	// Define the regex pattern
	// at least 8 characters
	if len(password) < 8 {
		return false
	}

	// at least one letter
	if matched, _ := regexp.MatchString("[a-zA-Z]", password); !matched {
		return false
	}

	// at least one number
	if matched, _ := regexp.MatchString("[0-9]", password); !matched {
		return false
	}

	// at least one special character @$!%*#?&
	if matched, _ := regexp.MatchString(`[@$!%*#?&]`, password); !matched {
		return false
	}

	return true
}
