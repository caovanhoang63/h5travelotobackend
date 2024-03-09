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
	if m, _ := regexp.MatchString(`^(?=.*[A-Za-z])(?=.*\d)(?=.*[@$!%*#?&])[A-Za-z\d@$!%*#?&]{8,}$`, password); !m {
		return false
	}
	return true
}
