package common

import (
	"github.com/asaskevich/govalidator"
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
