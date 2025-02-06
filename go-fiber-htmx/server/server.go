package server

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/jet/v2"
	"github.com/senither/custodian/config"
	"github.com/senither/custodian/server/handler"
	"github.com/senither/custodian/server/middleware"
	"github.com/senither/custodian/server/router"
)

func NewServer(cfg config.ServerConfig) *fiber.App {
	app := createNewFiberApp(cfg)

	middleware.Wrap(func(app *fiber.App) {
		router.RegisterRoutes(app)
	}, app, cfg)

	return app
}

func createNewFiberApp(cfg config.ServerConfig) *fiber.App {
	engine := jet.NewFileSystem(http.FS(cfg.ViewFilesystem), ".jet.html")

	return fiber.New(fiber.Config{
		AppName:      config.Get().Application.Name,
		ServerHeader: "Custodian Web Server",
		Views:        engine,
		ErrorHandler: handler.HandleInternalError,
	})
}
