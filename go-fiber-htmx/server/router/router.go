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

	hx.Get("/tasks", handler.RenderTasksComponent)
	hx.Get("/create-task-modal", handler.RenderCreateTaskModalComponent)
	hx.Get("/edit-task-modal/:task", handler.RenderEditTaskModalComponent)
	hx.Get("/delete-task-modal/:task", handler.RenderDeleteTaskModalComponent)
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
	}

	return func(c *fiber.Ctx) error {
		return c.Render(view, binds, layout)
	}
}
