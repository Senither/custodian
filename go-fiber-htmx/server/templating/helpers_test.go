package templating

import (
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
)

func TestGetInputErrorFor(t *testing.T) {
	t.Run("it returns formatted error messages", func(t *testing.T) {
		errors := &fiber.Map{
			"email": []string{"Email is required", "Email must be valid"},
		}

		expected := `<div class="mt-2"><ul class="text-error text-sm"><li>Email is required</li><li>Email must be valid</li></ul></div>`
		assert.Equal(t, expected, getInputErrorFor(errors, "email"))
	})

	t.Run("it returns empty string when no errors", func(t *testing.T) {
		errors := &fiber.Map{}

		assert.Equal(t, "", getInputErrorFor(errors, "email"))
	})

	t.Run("it returns empty string when errors is nil", func(t *testing.T) {
		assert.Equal(t, "", getInputErrorFor(nil, "email"))
	})
}

func TestGetOldInputValueFor(t *testing.T) {
	t.Run("it returns old input value", func(t *testing.T) {
		old := &fiber.Map{
			"email": "john@example.com",
		}

		assert.Equal(t, "john@example.com", getOldInputValueFor(old, "email"))
	})

	t.Run("it returns empty string when key does not exist", func(t *testing.T) {
		old := &fiber.Map{}

		assert.Equal(t, "", getOldInputValueFor(old, "email"))
	})

	t.Run("it returns empty string when old is nil", func(t *testing.T) {
		assert.Equal(t, "", getOldInputValueFor(nil, "email"))
	})
}
