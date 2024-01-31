package response

import (
	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	ent "github.com/go-playground/validator/v10/translations/en"
	"github.com/gofiber/fiber/v2"
	"github.com/rs/zerolog/log"
	"reflect"
)

var (
	validate *validator.Validate
	uni      *ut.UniversalTranslator
	Trans    ut.Translator
)

func init() {
	validate = validator.New()

	uni = ut.New(en.New())
	Trans, _ = uni.GetTranslator("en")

	if err := ent.RegisterDefaultTranslations(validate, Trans); err != nil && !fiber.IsChild() {
		log.Panic().Err(err).Msg("")
	}
}

func ValidateStruct(input any) error {
	return validate.Struct(input)
}

func ParseBody(c *fiber.Ctx, body any) error {
	if err := c.BodyParser(body); err != nil {
		return err
	}

	return nil
}

func ParseAndValidate(c *fiber.Ctx, body any) error {
	v := reflect.ValueOf(body)

	switch v.Kind() {
	case reflect.Ptr:
		err := ParseBody(c, body)
		if err != nil {
			return err
		}

		return ValidateStruct(v.Elem().Interface())
	case reflect.Struct:
		err := ParseBody(c, &body)
		if err != nil {
			return err
		}

		return ValidateStruct(v)
	default:
		return nil
	}
}
