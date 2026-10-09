package router

import (
	"errors"

	"github.com/fateevanastusha/blogs-go-backend/internal/model"
	"github.com/go-playground/validator/v10"
)

var validate = validator.New(validator.WithRequiredStructEnabled())

func validateStruct(s any) error {
	err := validate.Struct(s)
	if err == nil {
		return nil
	}
	var verrs validator.ValidationErrors
	if !errors.As(err, &verrs) {
		return err
	}
	ve := &model.ValidationError{}
	for _, fe := range verrs {
		ve.Fields = append(ve.Fields, model.FieldError{
			Field:   fe.Field(),
			Message: message(fe),
		})
	}
	return ve
}

func message(fe validator.FieldError) string {
	switch fe.Tag() {
	case "required":
		return "is required"
	case "max":
		return "must be a maximum of " + fe.Param() + " characters"
	case "url":
		return "must be a valid URL"
	case "startsWith":
		return "must start with " + fe.Param()
	}
	return "invalid"
}
