package server

import (
	"github.com/gofiber/fiber/v2"
	"github.com/senither/custodian/config"
	"github.com/senither/custodian/server/handler"
	"github.com/senither/custodian/server/middleware"
	"github.com/senither/custodian/server/router"
	"github.com/senither/custodian/server/session"
	"github.com/senither/custodian/server/templating"
)

func NewServer(cfg config.ServerConfig) *fiber.App {
	session.InitiateSessionStorage()

	app := createNewFiberApp(cfg)

	middleware.Wrap(func(app *fiber.App) {
		router.RegisterRoutes(app)
	}, app, cfg)

	return app
}

func createNewFiberApp(cfg config.ServerConfig) *fiber.App {
	engine := templating.NewTemplatingEngine(cfg)

	app := fiber.New(fiber.Config{
		AppName:      config.Get().Application.Name,
		ServerHeader: "Custodian Web Server",
		Views:        engine,
		ErrorHandler: handler.HandleInternalError,
	})

	templating.RegisterTemplateDefaults(engine)

	return app
}
