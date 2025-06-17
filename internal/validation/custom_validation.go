package validation

import (
	"main/internal/utils"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"

	"github.com/go-playground/validator/v10"
)

func RegisterCustomValidations(v *validator.Validate) {

	v.RegisterValidation("password_strong", func(fl validator.FieldLevel) bool {
		password := fl.Field().String()
		if len(password) < 8 {
			return false
		}
		hasUpper := false
		hasLower := false
		hasNumber := false
		hasSpecial := false

		for _, char := range password {
			switch {
			case char >= 'A' && char <= 'Z':
				hasUpper = true
			case char >= 'a' && char <= 'z':
				hasLower = true
			case char >= '0' && char <= '9':
				hasNumber = true
			case strings.ContainsRune("!@#$%^&*()-_=+[]{}|;:,.<>?/", char):
				hasSpecial = true
			}
		}

		return hasUpper && hasLower && hasNumber && hasSpecial
	})

	var blockedDomain = map[string]bool{
		"blacklist.com": true,
		"edu.com":       true,
		"abc.com":       true,
	}
	v.RegisterValidation("email_advanced", func(fl validator.FieldLevel) bool {
		email := fl.Field().String()
		parts := strings.Split(email, "@")
		if len(parts) != 2 {
			return false
		}
		domain := utils.NormalizeString(parts[1])
		return !blockedDomain[domain]

	})

	var slugRegex = regexp.MustCompile(`^[a-z0-9]+(?:-[a-z0-9]+)*$`)
	v.RegisterValidation("slug", func(fl validator.FieldLevel) bool {
		return slugRegex.MatchString(fl.Field().String())
	})

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
