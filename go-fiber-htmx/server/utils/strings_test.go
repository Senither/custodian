package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestToSnakeCase(t *testing.T) {
	t.Run("convert camelCase to snake_case", func(t *testing.T) {
		input := "camelCaseString"
		expected := "camel_case_string"

		assert.Equal(t, expected, ToSnakeCase(input))
	})

	t.Run("convert PascalCase to snake_case", func(t *testing.T) {
		input := "PascalCaseString"
		expected := "pascal_case_string"

		assert.Equal(t, expected, ToSnakeCase(input))
	})

	t.Run("handle already snake_case", func(t *testing.T) {
		input := "snake_case_string"
		expected := "snake_case_string"

		assert.Equal(t, expected, ToSnakeCase(input))
	})
}

func TestParseToUint(t *testing.T) {
	t.Run("it can parse valid numeric string", func(t *testing.T) {
		input := "123"
		expected := uint(123)

		assert.Equal(t, expected, ParseToUint(input))
	})

	t.Run("it returns 0 for invalid input", func(t *testing.T) {
		input := "invalid"
		expected := uint(0)

		assert.Equal(t, expected, ParseToUint(input))
	})

	t.Run("it defaults to 0 for empty strings", func(t *testing.T) {
		input := ""
		expected := uint(0)

		assert.Equal(t, expected, ParseToUint(input))
	})
}
