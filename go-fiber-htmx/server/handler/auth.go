package handler

import "github.com/gofiber/fiber/v2"

func RenderLogin(c *fiber.Ctx) error {
	return c.Render("views/auth/login", fiber.Map{}, "views/layouts/guest")
}

func RenderRegister(c *fiber.Ctx) error {
	return c.Render("views/auth/register", fiber.Map{}, "views/layouts/guest")
}

func RenderForgotPassword(c *fiber.Ctx) error {
	return c.Render("views/auth/forgot-password", fiber.Map{}, "views/layouts/guest")
}
