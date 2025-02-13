package handler

import (
	"log/slog"

	"github.com/gofiber/fiber/v2"
	"github.com/senither/custodian/database/model"
	"github.com/senither/custodian/database/repository"
	"github.com/senither/custodian/server/session"
	"github.com/senither/custodian/server/utils"
)

func RenderTasksComponent(c *fiber.Ctx) error {
	user, err := session.GetAuthenticatedUser(c)
	if err != nil {
		return c.SendString("Failed to load user from session")
	}

	return renderTaskComponentWithData(c, user)
}

func RenderFiltersComponent(c *fiber.Ctx) error {
	user, err := session.GetAuthenticatedUser(c)
	if err != nil {
		return c.SendString("Failed to load user from session")
	}

	categories, _ := repository.GetCategoriesForUser(c.UserContext(), user)
	priorities, _ := repository.GetPrioritiesForUser(c.UserContext(), user)

	return c.Render("views/components/filters", fiber.Map{
		"categories": categories,
		"priorities": priorities,
	})
}

func ToggleTaskStatus(c *fiber.Ctx) error {
	user, err := session.GetAuthenticatedUser(c)
	if err != nil {
		return c.SendString("Failed to load user from session")
	}

	taskId := utils.ParseToUint(c.Params("task"))
	task, dbErr := repository.FindTaskForUser(c.UserContext(), *user, taskId)
	if dbErr != nil {
		return renderTaskComponentWithData(c, user)
	}

	updateErr := repository.UpdateTask(c.UserContext(), *task, map[string]interface{}{
		"status": c.FormValue("status") == "on",
	})
	if updateErr != nil {
		slog.Error("Failed to update task status", "error", updateErr, "task", task)
	}

	return renderTaskComponentWithData(c, user)
}

func renderTaskComponentWithData(c *fiber.Ctx, user *model.User) error {
	search := buildSearchQuery(c)

	tasks, dbErr := repository.GetTasksWithSearchForUserWithRelations(c.UserContext(), user, search)
	if dbErr != nil {
		return c.Render("views/components/tasks", fiber.Map{
			"user":     user,
			"tasks":    tasks,
			"hasError": true,
		})
	}

	return c.Render("views/components/tasks", fiber.Map{
		"user":      user,
		"tasks":     tasks,
		"hasSearch": len(search) > 0,
	})
}

func buildSearchQuery(c *fiber.Ctx) map[string]interface{} {
	var search = make(map[string]interface{})

	if len(c.Query("q")) > 0 {
		search["message LIKE ?"] = "%" + c.Query("q") + "%"
	}

	if len(c.Query("category")) > 0 {
		search["category_id = ?"] = utils.ParseToUint(c.Query("category"))
	}

	if len(c.Query("priority")) > 0 {
		search["priority_id = ?"] = utils.ParseToUint(c.Query("priority"))
	}

	if len(c.Query("status")) > 0 {
		switch c.Query("status") {
		case "finished":
			search["status = ?"] = 1
		case "pending":
			search["status = ?"] = 0
		}
	}

	return search
}
