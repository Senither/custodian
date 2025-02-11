package handler

import (
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

	return c.SendStatus(fiber.StatusCreated)
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
