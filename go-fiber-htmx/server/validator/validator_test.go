package validator

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

type TestStruct struct {
	Name     string `validate:"required"`
	Email    string `validate:"required,email"`
	Age      int    `validate:"required,gte=18"`
	Password string `validate:"required,min=6"`
}

func TestParse(t *testing.T) {
	ctx := context.Background()

	t.Run("it passes when provided a struct with valid values", func(t *testing.T) {
		input := TestStruct{
			Name:     "John Doe",
			Email:    "john@example.com",
			Age:      25,
			Password: "secret123",
		}

		errors := Parse(ctx, input)
		assert.Nil(t, errors)
	})

	t.Run("it fails when provided a struct with invalid values", func(t *testing.T) {
		input := TestStruct{
			Name:     "",
			Email:    "invalid-email",
			Age:      16,
			Password: "123",
		}

		errors := Parse(ctx, input)
		assert.NotNil(t, errors)

		errMap := *errors
		assert.Len(t, errMap, 4)

		assert.Equal(t, errMap["name"].([]string)[0], "Name is a required field")
		assert.Equal(t, errMap["email"].([]string)[0], "Email must be a valid email address")
		assert.Equal(t, errMap["age"].([]string)[0], "Age must be 18 or greater")
		assert.Equal(t, errMap["password"].([]string)[0], "Password must be at least 6 characters in length")
	})

	t.Run("it fails when provided a struct with partial invalid values", func(t *testing.T) {
		input := TestStruct{
			Name:     "John Doe",
			Email:    "invalid-email",
			Age:      25,
			Password: "123",
		}

		errors := Parse(ctx, input)
		assert.NotNil(t, errors)

		errMap := *errors
		assert.Len(t, errMap, 2)

		assert.Equal(t, errMap["email"].([]string)[0], "Email must be a valid email address")
		assert.Equal(t, errMap["password"].([]string)[0], "Password must be at least 6 characters in length")
	})
}
