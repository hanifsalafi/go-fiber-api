package validator

import (
	"fmt"
	"github.com/go-playground/validator/v10"
	req "go-fiber-api/model/request"
)

func ValidateUserRequest(user *req.UserProfileCreateRequest) error {
	validate := validator.New()
	if err := validate.RegisterValidation("lowercase", LowercaseValidation); err != nil {
		return fmt.Errorf("error registering validator: %s", err)
	}

	if err := validate.Struct(user); err != nil {
		return fmt.Errorf("validator error: %s", err)
	}

	return nil
}
