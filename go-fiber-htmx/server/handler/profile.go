package handler

import "github.com/gofiber/fiber/v2"

func RenderProfile(c *fiber.Ctx) error {
	return c.Render("views/profile", fiber.Map{}, "views/layouts/app")
}
