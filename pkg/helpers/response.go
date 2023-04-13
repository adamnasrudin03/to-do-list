package helpers

import (
	"fmt"
	"strings"

	"github.com/go-playground/validator/v10"
)

type ResponseDefault struct {
	Status  string      `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

// APIResponse is for generating template responses
func APIResponse(message string, status string, data interface{}) ResponseDefault {

	return ResponseDefault{
		Status:  status,
		Message: message,
		Data:    data,
	}
}

// FormatValidationError func which holds errors during user input validation
func FormatValidationError(err error) string {
	var errors string

	for _, e := range err.(validator.ValidationErrors) {
		if errors != "" {
			errors = fmt.Sprintf("%v, ", strings.TrimSpace(errors))
		}

		if e.Tag() == "email" {
			errors = errors + fmt.Sprintf("%v must be type %v", e.Field(), e.Tag())
		} else {
			errors = errors + fmt.Sprintf("%v is %v %v", e.Field(), e.Tag(), e.Param())
		}

		if e.Param() != "" && e.Type().Name() == "string" {
			errors = errors + " character"
		}

	}

	return strings.TrimSpace(errors) + "."
}
