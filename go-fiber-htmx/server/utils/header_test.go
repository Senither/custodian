package utils

import (
	"io"
	"net/http/httptest"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
)

func TestGetRequestHeader(t *testing.T) {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(GetRequestHeader(c, "Test-Header"))
	})

	t.Run("it can get existing header", func(t *testing.T) {
		req := httptest.NewRequest("GET", "/", nil)
		req.Header.Set("Test-Header", "HeaderValue")
		resp, err := app.Test(req)

		assert.Nil(t, err)
		assert.Equal(t, 200, resp.StatusCode)

		expected := `["HeaderValue"]`
		body, _ := io.ReadAll(resp.Body)
		assert.JSONEq(t, expected, string(body))
	})

	t.Run("it returns empty array for non-existing header", func(t *testing.T) {
		req := httptest.NewRequest("GET", "/", nil)
		resp, err := app.Test(req)

		assert.Nil(t, err)
		assert.Equal(t, 200, resp.StatusCode)

		expected := `[]`
		body, _ := io.ReadAll(resp.Body)
		assert.JSONEq(t, expected, string(body))
	})
}
