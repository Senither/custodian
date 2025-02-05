package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/senither/custodian/config"
)

func HandleInternalError(c *fiber.Ctx, err error) error {
	code := fiber.StatusInternalServerError
	message := err.Error()

	if e, ok := err.(*fiber.Error); ok {
		code = e.Code
	}

	if !config.Get().Application.Debug {
		message = "An internal error occurred, please try again later."
	}

	return c.Status(code).Render("views/errors/internal", fiber.Map{
		"ErrorCode":             code,
		"ErrorMessage":          message,
		"ApplicationName":       config.Get().Application.Name,
		"ApplicationDescriptor": config.Get().Application.Descriptor,
	}, "views/layouts/guest")
}
