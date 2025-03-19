package utils

import (
	"io"
	"testing"

	"net/http/httptest"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
)

type TestStruct struct {
	Name  string
	Email string
	Age   int
}

func TestConvertToFiberMap(t *testing.T) {
	t.Run("it can convert struct to fiber map", func(t *testing.T) {
		input := TestStruct{
			Name:  "John Doe",
			Email: "john@example.com",
			Age:   30,
		}
		expected := &fiber.Map{
			"name":  "John Doe",
			"email": "john@example.com",
			"age":   30,
		}

		result := ConvertToFiberMap(input)
		assert.Equal(t, expected, result)
	})

	t.Run("it can convert map to fiber map", func(t *testing.T) {
		input := map[string]interface{}{
			"Name":  "John Doe",
			"Email": "john@example.com",
			"Age":   30,
		}
		expected := &fiber.Map{
			"name":  "John Doe",
			"email": "john@example.com",
			"age":   30,
		}

		result := ConvertToFiberMap(input)
		assert.Equal(t, expected, result)
	})

	t.Run("it can convert slice to fiber map", func(t *testing.T) {
		input := []string{"John Doe", "john@example.com", "30"}
		expected := &fiber.Map{
			"0": "John Doe",
			"1": "john@example.com",
			"2": "30",
		}

		result := ConvertToFiberMap(input)
		assert.Equal(t, expected, result)
	})
}

func TestSendHtmxRefreshScript(t *testing.T) {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return SendHtmxRefreshScript(c, "some-target")
	})

	t.Run("it sends the expected HTMX refresh script", func(t *testing.T) {
		req := httptest.NewRequest("GET", "/", nil)

		resp, err := app.Test(req)
		assert.Nil(t, err)

		expectedBody := "<script>window.htmx.trigger('some-target', 'refresh')</script>"
		body, _ := io.ReadAll(resp.Body)
		assert.Equal(t, expectedBody, string(body))
	})
}
