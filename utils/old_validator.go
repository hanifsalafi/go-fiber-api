package utils

import (
	"fmt"
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

func ValidateRequest(any interface{}) error {
	validate := validator.New()
	if err := validate.RegisterValidation("lowercase", LowercaseValidation); err != nil {
		return fmt.Errorf("error registering validator: %s", err)
	}

	if err := validate.Struct(any); err != nil {
		return fmt.Errorf("%s", err)
	}

	return nil
}
