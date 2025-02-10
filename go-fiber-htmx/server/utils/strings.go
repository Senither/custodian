package utils

import "strings"

func ToSnakeCase(s string) string {
	return strings.ToLower(snakeCaseRegexp.ReplaceAllString(s, "${1}_${2}"))
}
