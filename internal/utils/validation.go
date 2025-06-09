package utils

import (
	"fmt"
	"regexp"
)

func ValidationRequired(fieldName, value string) error {
	if value == "" {
		return fmt.Errorf("%s is required", fieldName)
	}
	return nil
}

func ValidationStringLength(fieldName, value string, max, min int) error {
	l := len(value)
	if l < min {
		return fmt.Errorf("%s is too short", fieldName)
	}
	if l > max {
		return fmt.Errorf("%s is too long", fieldName)
	}
	return nil
}

func ValidationRegex(fieldName, value string, regex *regexp.Regexp) error {

	if !regex.MatchString(value) {
		return fmt.Errorf("%s is invalid", fieldName)
	}
	return nil
}
