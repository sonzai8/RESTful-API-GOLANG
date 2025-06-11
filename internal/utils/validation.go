package utils

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	"log"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
)

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
					base := camelToSnake(p[:idx])
					index := p[idx:]
					parts[i] = base + index
				} else {
					parts[i] = camelToSnake(p)
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

func RegisterValidators() error {
	v, ok := binding.Validator.Engine().(*validator.Validate)
	if !ok {
		return fmt.Errorf("validate validator engine failed")
	}

	var slugRegex = regexp.MustCompile(`^[a-z0-9]+(?:-[a-z0-9]+)*$`)
	err := v.RegisterValidation("slug", func(fl validator.FieldLevel) bool {
		return slugRegex.MatchString(fl.Field().String())
	})
	if err != nil {
		panic(err)
	}

	var searchRegex = regexp.MustCompile(`^[a-zA-Z0-9\s]+$`)
	v.RegisterValidation("search", func(fl validator.FieldLevel) bool {
		return searchRegex.MatchString(fl.Field().String())
	})

	//min_int:1000
	v.RegisterValidation("min_int", func(fl validator.FieldLevel) bool {
		minStr := fl.Param()
		minVal, err := strconv.ParseInt(minStr, 10, 64)
		if err != nil {
			return false
		}
		return fl.Field().Int() >= minVal
	})

	v.RegisterValidation("max_int", func(fl validator.FieldLevel) bool {
		maxStr := fl.Param()
		maxVal, err := strconv.ParseInt(maxStr, 10, 64)
		if err != nil {
			return false
		}
		return fl.Field().Int() <= maxVal
	})

	v.RegisterValidation("file_ext", func(fl validator.FieldLevel) bool {
		filename := fl.Field().String()

		allowStr := fl.Param()

		if allowStr == "" {
			return false
		}

		allowExt := strings.Fields(allowStr)

		ext := strings.TrimPrefix(strings.ToLower(filepath.Ext(filename)), ".")

		for _, allowed := range allowExt {
			if strings.ToLower(allowed) == ext {
				return true
			}
		}
		return false
	})

	return nil
}
