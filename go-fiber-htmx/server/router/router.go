package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
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
	app.Get("/dashboard", handler.RenderDashboard)
	app.Get("/profile", handler.RenderProfile)

	// Guest routes
	app.Get("/login", handler.RenderLogin)
	app.Get("/forgot-password", handler.RenderForgotPassword)
	app.Get("/register", handler.RenderRegister)
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
