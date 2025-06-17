package utils

import (
	"regexp"
	"strings"
)

var (
	matchFirstCap = regexp.MustCompile("(.)([A-Z][a-z]+)")
	matchAllCap   = regexp.MustCompile("([a-z0-9])([A-Z])")
)

func CamelToSnake(str string) string {
	snake := matchFirstCap.ReplaceAllString(str, "${1}_${2}")
	snake = matchAllCap.ReplaceAllString(snake, "${1}_${2}")
	return strings.ToLower(snake)
}

func NormalizeString(str string) string {
	// Convert to lowercase
	str = strings.ToLower(str)
	str = strings.TrimSpace(str)
	// Replace spaces with underscores
	// str = strings.ReplaceAll(str, " ", "_")
	// Remove special characters except underscores and alphanumeric characters
	// str = regexp.MustCompile(`[^a-z0-9_]+`).ReplaceAllString(str, "")
	return str
}

// 			errors[fieldPath] = e.Tag()
