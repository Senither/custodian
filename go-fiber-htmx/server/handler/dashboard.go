package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/senither/custodian/database/repository"
	"github.com/senither/custodian/server/session"
)

func RenderTasksComponent(c *fiber.Ctx) error {
	user, err := session.GetAuthenticatedUser(c)
	if err != nil {
		return c.SendString("Failed to load user from session")
	}

	return c.Render("views/components/tasks", fiber.Map{
		"user": user,
	})
}

func RenderCreateTaskModalComponent(c *fiber.Ctx) error {
	user, err := session.GetAuthenticatedUser(c)
	if err != nil {
		return c.SendString("Failed to load user from session")
	}

	categories, _ := repository.GetCategoriesForUser(c.UserContext(), user)
	priorities, _ := repository.GetPrioritiesForUser(c.UserContext(), user)

	return c.Render("views/components/create-task-modal", fiber.Map{
		"categories": categories,
		"priorities": priorities,
	})
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
