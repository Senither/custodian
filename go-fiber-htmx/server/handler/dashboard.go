package handler

import "github.com/gofiber/fiber/v2"

func RenderTasksComponent(c *fiber.Ctx) error {
	return c.Render("views/components/tasks", fiber.Map{})
}
