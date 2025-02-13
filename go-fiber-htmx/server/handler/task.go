package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/senither/custodian/database/model"
	"github.com/senither/custodian/database/repository"
	"github.com/senither/custodian/server/session"
	"github.com/senither/custodian/server/utils"
	"github.com/senither/custodian/server/validator"
)

type CreateOrUpdateTaskRequest struct {
	Message    string `validate:"required,min=1,max=255"`
	CategoryId uint   `validate:"required,gte=1"`
	PriorityId uint   `validate:"required,gte=1"`
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

func CreateTask(c *fiber.Ctx) error {
	user, err := session.GetAuthenticatedUser(c)
	if err != nil {
		return c.SendString("Failed to load user from session")
	}

	categories, _ := repository.GetCategoriesForUser(c.UserContext(), user)
	priorities, _ := repository.GetPrioritiesForUser(c.UserContext(), user)

	request := CreateOrUpdateTaskRequest{
		Message:    c.FormValue("message"),
		CategoryId: utils.ParseToUint(c.FormValue("category_id")),
		PriorityId: utils.ParseToUint(c.FormValue("priority_id")),
	}

	category, priority, valErr := validateCreateOrUpdateTaskRequest(c, categories, priorities, request)
	if valErr != nil {
		return c.Render("views/components/create-task-modal", fiber.Map{
			"old":        utils.ConvertToFiberMap(request),
			"errors":     valErr,
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

	return utils.SendHtmxRefreshScript(c, "#tasks")
}

func RenderEditTaskModalComponent(c *fiber.Ctx) error {
	user, err := session.GetAuthenticatedUser(c)
	if err != nil {
		return c.SendString("Failed to load user from session")
	}

	task, dbErr := repository.FindTaskForUser(c.UserContext(), *user, utils.ParseToUint(c.Params("task")))
	if dbErr != nil {
		return c.SendString("Failed to load task")
	}

	categories, _ := repository.GetCategoriesForUser(c.UserContext(), user)
	priorities, _ := repository.GetPrioritiesForUser(c.UserContext(), user)

	return c.Render("views/components/edit-task-modal", fiber.Map{
		"task":       task,
		"categories": categories,
		"priorities": priorities,
	})
}

func UpdateTask(c *fiber.Ctx) error {
	user, err := session.GetAuthenticatedUser(c)
	if err != nil {
		return c.SendString("Failed to load user from session")
	}

	task, dbErr := repository.FindTaskForUser(c.UserContext(), *user, utils.ParseToUint(c.Params("task")))
	if dbErr != nil {
		return c.SendString("Failed to load task")
	}

	categories, _ := repository.GetCategoriesForUser(c.UserContext(), user)
	priorities, _ := repository.GetPrioritiesForUser(c.UserContext(), user)

	request := CreateOrUpdateTaskRequest{
		Message:    c.FormValue("message"),
		CategoryId: utils.ParseToUint(c.FormValue("category_id")),
		PriorityId: utils.ParseToUint(c.FormValue("priority_id")),
	}

	category, priority, valErr := validateCreateOrUpdateTaskRequest(c, categories, priorities, request)
	if valErr != nil {
		return c.Render("views/components/edit-task-modal", fiber.Map{
			"task":       task,
			"errors":     valErr,
			"categories": categories,
			"priorities": priorities,
		})
	}

	updateErr := repository.UpdateTask(c.UserContext(), *task, map[string]interface{}{
		"message":     request.Message,
		"category_id": category.ID,
		"priority_id": priority.ID,
	})

	if updateErr != nil {
		return c.Render("views/components/edit-task-modal", fiber.Map{
			"task": task,
			"errors": &fiber.Map{
				"message": []string{"Failed to save the task, please try again later"},
			},
			"categories": categories,
			"priorities": priorities,
		})
	}

	return utils.SendHtmxRefreshScript(c, "#tasks")
}

func validateCreateOrUpdateTaskRequest(
	c *fiber.Ctx,
	categories []model.Category,
	priorities []model.Priority,
	request CreateOrUpdateTaskRequest,
) (*model.Category, *model.Priority, *fiber.Map) {
	if err := validator.Parse(c.UserContext(), request); err != nil {
		return nil, nil, err
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
		return nil, nil, &errors
	}
	return category, priority, nil
}

func RenderDeleteTaskModalComponent(c *fiber.Ctx) error {
	user, err := session.GetAuthenticatedUser(c)
	if err != nil {
		return c.SendString("Failed to load user from session")
	}

	task, dbErr := repository.FindTaskForUser(c.UserContext(), *user, utils.ParseToUint(c.Params("task")))
	if dbErr != nil {
		return c.SendString("Failed to load task")
	}

	return c.Render("views/components/delete-task-modal", fiber.Map{
		"task": task,
	})
}

func DeleteTask(c *fiber.Ctx) error {
	user, err := session.GetAuthenticatedUser(c)
	if err != nil {
		return c.SendString("Failed to load user from session")
	}

	task, dbErr := repository.FindTaskForUser(c.UserContext(), *user, utils.ParseToUint(c.Params("task")))
	if dbErr != nil {
		return c.SendString("Failed to load task")
	}

	deleteErr := repository.DeleteTask(c.UserContext(), *task)
	if deleteErr != nil {
		return c.SendString("Failed to delete task")
	}

	return utils.SendHtmxRefreshScript(c, "#tasks")
}
