package router

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/senither/custodian/config"
	"github.com/senither/custodian/server/handler"
	"github.com/senither/custodian/server/middleware"
	"github.com/senither/custodian/server/session"
)

func RegisterRoutes(app *fiber.App) {
	registerViewRoutes(app)
	registerHtmxRoutes(app)
	registerRedirectRoutes(app)
}

func registerViewRoutes(app *fiber.App) {
	// Authenticated routes
	app.Get("/dashboard", middleware.Authenticated(), createViewWithLayoutHandler("dashboard", "app")).Name("dashboard")
	app.Get("/profile", middleware.Authenticated(), createViewWithLayoutHandler("profile", "app")).Name("profile")
	app.Get("/logout", middleware.Authenticated(), handler.Logout).Name("logout")

	// Guest routes
	app.Get("/login", middleware.Guest(), createViewWithLayoutHandler("auth/login", "guest")).Name("login")
	app.Post("/login", middleware.Guest(), handler.Login)

	app.Get("/register", middleware.Guest(), createViewWithLayoutHandler("auth/register", "guest")).Name("register")
	app.Post("/register", middleware.Guest(), handler.Register)

	app.Get("/forgot-password", middleware.Guest(), createViewWithLayoutHandler("auth/forgot-password", "guest")).Name("forgot-password")
	app.Post("/forgot-password", middleware.Guest(), handler.ForgotPassword)
}

func registerHtmxRoutes(app *fiber.App) {
	hx := app.Group("/hx")

	// Tasks
	hx.Get("/tasks", handler.RenderTasksComponent)
	hx.Post("/toggle-task-status/:task", handler.ToggleTaskStatus)
	hx.Get("/create-task-modal", handler.RenderCreateTaskModalComponent)
	hx.Post("/create-task-modal", handler.CreateTask)
	hx.Get("/edit-task-modal/:task", handler.RenderEditTaskModalComponent)
	hx.Post("/edit-task-modal/:task", handler.UpdateTask)
	hx.Get("/delete-task-modal/:task", handler.RenderDeleteTaskModalComponent)
	hx.Post("/delete-task-modal/:task", handler.DeleteTask)

	// Profile
	hx.Post("/update-profile-information", handler.UpdateProfileInformation)
	hx.Post("/update-profile-password", handler.UpdateProfilePassword)
	hx.Get("/delete-account-modal", handler.RenderDeleteAccountModalComponent)
	hx.Post("/delete-account-modal", handler.DeleteAccount)
}

func registerRedirectRoutes(app *fiber.App) {
	app.Get("/", func(c *fiber.Ctx) error {
		_, err := session.GetAuthenticatedUser(c)
		if err != nil {
			return c.Redirect("/login")
		}

		return c.Redirect("/dashboard")
	})
}

func createViewWithLayoutHandler(view string, layout string) func(*fiber.Ctx) error {
	view = fmt.Sprintf("views/%s", view)
	layout = fmt.Sprintf("views/layouts/%s", layout)

	binds := fiber.Map{
		"ApplicationName":       config.Get().Application.Name,
		"ApplicationDescriptor": config.Get().Application.Descriptor,
		"AuthenticatedUser":     fiber.Map{},
	}

	return func(c *fiber.Ctx) error {
		user, _ := session.GetAuthenticatedUserWithoutRedirects(c)
		if user != nil {
			binds["AuthenticatedUser"] = user
		}

		return c.Render(view, binds, layout)
	}
}
