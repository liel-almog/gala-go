package configs

import (
	"sync"

	"github.com/go-playground/validator/v10"
)

var (
	initValidatorOnce sync.Once
	validate          *validator.Validate
)

func GetValidator() *validator.Validate {
	initValidatorOnce.Do(func() {
		validate = validator.New()

		// Register custom validator
		validate.RegisterValidation("password", func(fl validator.FieldLevel) bool {
			return false
		})
	})

	return validate
}
