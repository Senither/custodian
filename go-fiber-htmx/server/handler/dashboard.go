package handler

import (
	"log/slog"

	"github.com/gofiber/fiber/v2"
	"github.com/senither/custodian/server/session"
)

func RenderTasksComponent(c *fiber.Ctx) error {
	user, err := session.GetAuthenticatedUser(c)
	if err != nil {
		slog.Error("Failed to load user from session", "err", err)

		return c.SendString("Failed to load user from session")
	}

	return c.Render("views/components/tasks", fiber.Map{
		"user": user,
	})
}

func RenderCreateTaskModalComponent(c *fiber.Ctx) error {
	return c.Render("views/components/create-task-modal", nil)
}

func RenderEditTaskModalComponent(c *fiber.Ctx) error {
	return c.Render("views/components/edit-task-modal", fiber.Map{
		"task": c.Params("task"),
	})
}

func RenderDeleteTaskModalComponent(c *fiber.Ctx) error {
	return c.Render("views/components/delete-task-modal", fiber.Map{
		"task": c.Params("task"),
	})
}
