package helper

import "github.com/go-playground/validator/v10"

var validate *validator.Validate

func ValidateJSON(data interface{}) (bool, map[string]string) {
	validate = validator.New()

	if err := validate.Struct(data); err != nil {
		errors := make(map[string]string)

		for _, err := range err.(validator.ValidationErrors) {
			errors[err.Field()] = err.Tag()
		}

		return false, errors
	}

	return true, nil
}
