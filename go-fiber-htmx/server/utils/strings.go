package utils

import (
	"strconv"
	"strings"
)

func ToSnakeCase(s string) string {
	return strings.ToLower(snakeCaseRegexp.ReplaceAllString(s, "${1}_${2}"))
}

func ParseToUint(value string) uint {
	parsedValue, err := strconv.ParseUint(value, 10, 32)
	if err != nil {
		return 0
	}

	return uint(parsedValue)
}
