package router

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/senither/custodian/config"
	"github.com/senither/custodian/server/handler"
)

func RegisterRoutes(app *fiber.App) {
	app.Use(logger.New())

	registerViewRoutes(app)
	registerHtmxRoutes(app)
	registerRedirectRoutes(app)
}

func registerViewRoutes(app *fiber.App) {
	// Authenticated routes
	app.Get("/dashboard", createViewWithLayoutHandler("dashboard", "app"))
	app.Get("/profile", createViewWithLayoutHandler("profile", "app"))

	// Guest routes
	app.Get("/login", createViewWithLayoutHandler("auth/login", "guest"))
	app.Get("/forgot-password", createViewWithLayoutHandler("auth/forgot-password", "guest"))
	app.Get("/register", createViewWithLayoutHandler("auth/register", "guest"))
}

func registerHtmxRoutes(app *fiber.App) {
	hx := app.Group("/hx")

	hx.Get("/tasks", handler.RenderTasksComponent)
}

func registerRedirectRoutes(app *fiber.App) {
	app.Get("/", func(c *fiber.Ctx) error {
		// Redirect the user to the dashboard if they are already logged in
		// otherwise redirect them to the login page.
		// return c.Redirect("/login")

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
