package util

import (
	"fmt"
	"github.com/go-playground/validator/v10"
	"os"
)

const (
	ValidationError            = "validation error"
)

type ErrorResponse struct {
	Rule         	string	`json:"rule"`
	Value       	string	`json:"value"`
	Field       	string	`json:"field"`
}

type Model interface {
	ModelName() string
}

type ErrorWrapper struct {
	Type string
	Details interface{}
}

var validate = validator.New()

func ProjectPath() string {
	WD, _ := os.Getwd()

	return WD
}

func SeederPath() string  {
	return fmt.Sprintf("%s/database/seeders", ProjectPath())
}

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
