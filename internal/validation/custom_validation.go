package validation

import (
	"path/filepath"
	"regexp"
	"strconv"
	"strings"

	"github.com/go-playground/validator/v10"
)

func RegisterCustomValidations(v *validator.Validate) {

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

}
