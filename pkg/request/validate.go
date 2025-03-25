package request

import (
	"github.com/go-playground/validator/v10"
)

func Valid[T any](payload T) error {
	validate := validator.New()
	return validate.Struct(payload)
}
