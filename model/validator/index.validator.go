package validator

import (
	"github.com/go-playground/validator/v10"
	"unicode"
)

func LowercaseValidation(fl validator.FieldLevel) bool {
	value, _ := fl.Field().Interface().(string)

	for _, char := range value {
		if !unicode.IsLower(char) {
			return false
		}
	}

	return true
}
