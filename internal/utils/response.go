package utils

import (
	"github.com/go-playground/validator/v10"
)

func FormatValidationError(err error) []string {
	var errors []string
	if ve, ok := err.(validator.ValidationErrors); ok {
		for _, fe := range ve {
			errors = append(errors, formatError(fe))
		}
	} else {
		errors = append(errors, err.Error())
	}
	return errors
}

func formatError(fe validator.FieldError) string {
	switch fe.Tag() {
	case "required":
		return fe.Field() + " is required"
	case "email":
		return fe.Field() + " must be a valid email"
	case "min":
		return fe.Field() + " must be at least " + fe.Param() + " characters"
	}
	return fe.Field() + " is invalid"
}
