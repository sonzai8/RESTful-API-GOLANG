package validation

import (
	"fmt"
	"log"
	"main/internal/utils"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

func InitValidator() error {

	v, ok := binding.Validator.Engine().(*validator.Validate)
	if !ok {
		return fmt.Errorf("validate validator engine failed")
	}

	RegisterCustomValidations(v)
	return nil

}

func ValidationPositiveInt(fieldName, value string) (int, error) {
	v, err := strconv.Atoi(value)
	if err != nil {
		return 0, fmt.Errorf("%s is must be number", fieldName)
	}
	if v < 0 {
		return 0, fmt.Errorf("%s is must be positive", fieldName)
	}
	return v, nil
}

func HandleValidationErrors(err error) gin.H {
	if validationError, ok := err.(validator.ValidationErrors); ok {
		errors := make(map[string]string)

		for _, e := range validationError {
			log.Printf("%s", e.Namespace())
			root := strings.Split(e.Namespace(), ".")[0]
			rawPath := strings.TrimPrefix(e.Namespace(), root+".")

			parts := strings.Split(rawPath, ".")

			for i, p := range parts {
				if strings.Contains(p, "[") {
					idx := strings.Index(parts[i], "[")
					base := utils.CamelToSnake(p[:idx])
					index := p[idx:]
					parts[i] = base + index
				} else {
					parts[i] = utils.CamelToSnake(p)
				}

			}
			fieldPath := strings.Join(parts, ".")

			switch e.Tag() {
			case "required":
				errors[fieldPath] = fmt.Sprintf("%s is required", fieldPath)
			case "gt":
				errors[fieldPath] = fmt.Sprintf("%s must be greater than %s", fieldPath, e.Param())
			case "lt":
				errors[fieldPath] = fmt.Sprintf("%s must be less than %s", fieldPath, e.Param())
			case "slug":
				errors[fieldPath] = fmt.Sprintf("%s is must be slug", fieldPath)
			case "min":
				errors[fieldPath] = fmt.Sprintf("%s must be greater than %s", fieldPath, e.Param())
			case "max":
				errors[fieldPath] = fmt.Sprintf("%s must be less than %s", fieldPath, e.Param())
			case "oneof":
				errors[fieldPath] = fmt.Sprintf("%s must be one of the following values: %s",
					fieldPath, strings.Join(strings.Split(e.Param(), " "), ","))
			case "min_int":
				errors[fieldPath] = fmt.Sprintf("%s must be less than %s", fieldPath, e.Param())
			case "max_int":
				errors[fieldPath] = fmt.Sprintf("%s must be greater than %s", fieldPath, e.Param())
			case "file_ext":
				errors[fieldPath] = fmt.Sprintf("%s invalid ext: %s", fieldPath, e.Param())
			case "email":
				errors[fieldPath] = fmt.Sprintf("%s must be in correct email format: %s", fieldPath, e.Param())

			default:
				errors[fieldPath] = fmt.Sprintf("%s is invalid", fieldPath)
			}

		}
		return gin.H{"errors": errors}
	}
	return gin.H{"error": "yêu cầu không hợp lệ " + err.Error()}
}
