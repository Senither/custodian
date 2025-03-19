package utils

import (
	"github.com/gofiber/fiber/v2"
)

func RedirectWithHtmx(c *fiber.Ctx, url string) error {
	val := GetRequestHeader(c, "Hx-Request")

	if len(val) != 1 || val[0] != "true" {
		return c.Redirect(url)
	}

	c.Append("HX-Redirect", url)

	return c.SendString("HX Redirect to " + url)
}
