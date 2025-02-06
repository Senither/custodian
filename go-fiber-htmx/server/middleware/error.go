package middleware

import (
	"github.com/gofiber/fiber/v2"
	"github.com/senither/custodian/config"
)

func handlePageNotFound(c *fiber.Ctx) error {
	return c.
		Status(fiber.StatusNotFound).
		Render("views/errors/404", fiber.Map{
			"ApplicationName":       config.Get().Application.Name,
			"ApplicationDescriptor": config.Get().Application.Descriptor,
		}, "views/layouts/guest")
}
