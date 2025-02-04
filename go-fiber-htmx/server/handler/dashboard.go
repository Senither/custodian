package handler

import "github.com/gofiber/fiber/v2"

func RenderDashboard(c *fiber.Ctx) error {
	return c.Render("views/dashboard", fiber.Map{}, "views/layouts/app")
}

func RenderTasksComponent(c *fiber.Ctx) error {
	return c.Render("views/components/tasks", fiber.Map{})
}
