package handler

import (
	"log/slog"

	"github.com/gofiber/fiber/v2"
	"github.com/senither/custodian/database/model"
	"github.com/senither/custodian/database/repository"
	"github.com/senither/custodian/server/session"
	"github.com/senither/custodian/server/utils"
	"github.com/senither/custodian/server/validator"
)

func RenderTasksComponent(c *fiber.Ctx) error {
	user, err := session.GetAuthenticatedUser(c)
	if err != nil {
		return c.SendString("Failed to load user from session")
	}

	return renderTasksComponent(c, user)
}

func ToggleTaskStatus(c *fiber.Ctx) error {
	user, err := session.GetAuthenticatedUser(c)
	if err != nil {
		return c.SendString("Failed to load user from session")
	}

	taskId := utils.ParseToUint(c.Params("task"))
	task, dbErr := repository.FindTaskForUser(c.UserContext(), *user, taskId)
	if dbErr != nil {
		return renderTasksComponent(c, user)
	}

	updateErr := repository.UpdateTask(c.UserContext(), *task, map[string]interface{}{
		"status": c.FormValue("status") == "on",
	})
	if updateErr != nil {
		slog.Error("Failed to update task status", "error", updateErr, "task", task)
	}

	return renderTasksComponent(c, user)
}

func renderTasksComponent(c *fiber.Ctx, user *model.User) error {
	tasks, dbErr := repository.GetTasksForUserWithRelations(c.UserContext(), user)
	if dbErr != nil {
		return c.Render("views/components/tasks", fiber.Map{
			"user":     user,
			"tasks":    tasks,
			"hasError": true,
		})
	}

	return c.Render("views/components/tasks", fiber.Map{
		"user":  user,
		"tasks": tasks,
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

type CreateTaskRequest struct {
	Message    string `validate:"required,min=1,max=255"`
	CategoryId uint   `validate:"required,gte=1"`
	PriorityId uint   `validate:"required,gte=1"`
}

func CreateTask(c *fiber.Ctx) error {
	user, err := session.GetAuthenticatedUser(c)
	if err != nil {
		return c.SendString("Failed to load user from session")
	}

	categories, _ := repository.GetCategoriesForUser(c.UserContext(), user)
	priorities, _ := repository.GetPrioritiesForUser(c.UserContext(), user)

	request := CreateTaskRequest{
		Message:    c.FormValue("message"),
		CategoryId: utils.ParseToUint(c.FormValue("category_id")),
		PriorityId: utils.ParseToUint(c.FormValue("priority_id")),
	}

	if err := validator.Parse(c.UserContext(), request); err != nil {
		return c.Render("views/components/create-task-modal", fiber.Map{
			"old":        utils.ConvertToFiberMap(request),
			"errors":     err,
			"categories": categories,
			"priorities": priorities,
		})
	}

	var category *model.Category = nil
	var priority *model.Priority = nil
	errors := make(fiber.Map)

	for _, c := range categories {
		if c.ID == request.CategoryId {
			category = &c
			break
		}
	}

	if category == nil {
		errors["category_id"] = []string{"The selected category does not exist"}
	}

	for _, p := range priorities {
		if p.ID == request.PriorityId {
			priority = &p
			break
		}
	}

	if priority == nil {
		errors["priority_id"] = []string{"The selected priority does not exist"}
	}

	if len(errors) > 0 {
		return c.Render("views/components/create-task-modal", fiber.Map{
			"old":        utils.ConvertToFiberMap(request),
			"errors":     &errors,
			"categories": categories,
			"priorities": priorities,
		})
	}

	dbErr := repository.CreateTaskForUser(c.UserContext(), user, model.Task{
		Message:    request.Message,
		CategoryId: category.ID,
		PriorityId: priority.ID,
	})

	if dbErr != nil {
		return c.Render("views/components/create-task-modal", fiber.Map{
			"old": utils.ConvertToFiberMap(request),
			"errors": &fiber.Map{
				"message": []string{"Failed to create the task, please try again later"},
			},
			"categories": categories,
			"priorities": priorities,
		})
	}

	// Refresh the tasks component to show the new task
	return c.SendString("<script>window.htmx.trigger('#tasks', 'refresh')</script>")
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
