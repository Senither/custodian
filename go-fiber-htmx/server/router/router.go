package router

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/senither/custodian/config"
	"github.com/senither/custodian/server/handler"
	"github.com/senither/custodian/server/session"
)

func RegisterRoutes(app *fiber.App) {
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
	app.Post("/login", handler.Login)
	app.Get("/logout", handler.Logout)

	app.Get("/register", createViewWithLayoutHandler("auth/register", "guest"))
	app.Post("/register", handler.Register)

	app.Get("/forgot-password", createViewWithLayoutHandler("auth/forgot-password", "guest"))
	app.Post("/forgot-password", handler.ForgotPassword)
}

func registerHtmxRoutes(app *fiber.App) {
	hx := app.Group("/hx")

	hx.Get("/tasks", handler.RenderTasksComponent)
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
