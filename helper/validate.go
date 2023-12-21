package helper

import (
	"mime/multipart"
	"unicode"

	"github.com/go-playground/validator/v10"
)

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

func ValidateForm(data interface{}) (bool, map[string]string) {
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

func ValidateFile(file *multipart.FileHeader, maxSize int64, allowedTypes ...string) (bool, string) {
	if file == nil {
		return false, "File is required"
	}

	if file.Size > maxSize {
		return false, "File size exceeds the allowed limit"
	}

	allowed := false
	for _, allowedType := range allowedTypes {
		if file.Header.Get("Content-Type") == allowedType {
			allowed = true
			break
		}
	}

	if !allowed {
		return false, "Invalid file type"
	}

	return true, ""
}

func PasswordWithCombination(password string) bool {
	hasLetter := false
	hasDigit := false
	hasSymbol := false

	for _, char := range password {
		if unicode.IsLetter(char) {
			hasLetter = true
		} else if unicode.IsDigit(char) {
			hasDigit = true
		} else if unicode.IsPunct(char) || unicode.IsSymbol(char) {
			hasSymbol = true
		}
	}

	return hasLetter && hasDigit && hasSymbol
}
