package middleware

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/filesystem"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/senither/custodian/config"
)

func Wrap(wrapHandler func(*fiber.App), app *fiber.App, cfg config.ServerConfig) {
	RegisterBeforeMiddleware(app, cfg)

	wrapHandler(app)

	RegisterAfterMiddleware(app, cfg)
}

func RegisterBeforeMiddleware(app *fiber.App, cfg config.ServerConfig) {
	app.Use("/public", filesystem.New(filesystem.Config{
		Root:       http.FS(cfg.PublicFilesystem),
		PathPrefix: "/public",
	}))

	app.Use(newFiberLogger())
	app.Use(recover.New(recover.Config{
		EnableStackTrace: config.Get().Application.Debug,
	}))
	app.Use(handleSessions)
}

func RegisterAfterMiddleware(app *fiber.App, cfg config.ServerConfig) {
	app.Use(handlePageNotFound)
}
