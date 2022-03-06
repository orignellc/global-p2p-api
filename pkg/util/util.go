package util

import (
	"github.com/go-playground/validator/v10"
)

type ErrorResponse struct {
	Rule         	string	`json:"rule"`
	Value       	string	`json:"value"`
	Field       	string	`json:"field"`
}

type Model interface {
	ModelName() string
}

var validate = validator.New()

func ValidateStruct(model Model) []*ErrorResponse {
	var errors []*ErrorResponse
	err := validate.Struct(model)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			var element ErrorResponse
			element.Rule = err.Tag()
			element.Value = err.Param()
			element.Field = err.Field()
			errors = append(errors, &element)
		}
	}
	return errors
}
