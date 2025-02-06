package middleware

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/filesystem"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/senither/custodian/config"
)

func Wrap(wrapHandler func(*fiber.App), app *fiber.App, cfg config.ServerConfig) {
	RegisterBeforeMiddleware(app, cfg)

	wrapHandler(app)

	RegisterAfterMiddleware(app, cfg)
}

func RegisterBeforeMiddleware(app *fiber.App, cfg config.ServerConfig) {
	app.Use(logger.New())
	app.Use(handleSessions)
}

func RegisterAfterMiddleware(app *fiber.App, cfg config.ServerConfig) {
	app.Use("", filesystem.New(filesystem.Config{
		Root: http.FS(cfg.PublicFilesystem),
	}))

	app.Use(handlePageNotFound)
}
