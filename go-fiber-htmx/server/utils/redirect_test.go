package utils

import (
	"io"
	"net/http/httptest"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
)

func TestRedirectWithHtmx(t *testing.T) {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return RedirectWithHtmx(c, "/new-url")
	})

	t.Run("it redirects without Hx-Request header for non-HTMX requests", func(t *testing.T) {
		req := httptest.NewRequest("GET", "/", nil)
		resp, err := app.Test(req)

		assert.Nil(t, err)
		assert.Equal(t, 302, resp.StatusCode)
		assert.Equal(t, "/new-url", resp.Header.Get("Location"))
	})

	t.Run("it redirects with Hx-Request header for HTMX requests", func(t *testing.T) {
		req := httptest.NewRequest("GET", "/", nil)
		req.Header.Set("Hx-Request", "true")
		resp, err := app.Test(req)

		assert.Nil(t, err)
		assert.Equal(t, 200, resp.StatusCode)
		assert.Equal(t, "/new-url", resp.Header.Get("HX-Redirect"))

		expectedBody := "HX Redirect to /new-url"
		body, _ := io.ReadAll(resp.Body)
		assert.Equal(t, expectedBody, string(body))
	})
}
