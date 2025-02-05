package server

import (
	"embed"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/filesystem"
	"github.com/gofiber/template/jet/v2"
	"github.com/senither/custodian/config"
	"github.com/senither/custodian/server/router"
)

type ServerConfig struct {
	PublicFilesystem embed.FS
	ViewFilesystem   embed.FS
}

func NewServer(cfg ServerConfig) *fiber.App {
	app := createNewFiberApp(cfg)

	router.RegisterRoutes(app)

	app.Use("", filesystem.New(filesystem.Config{
		Root: http.FS(cfg.PublicFilesystem),
	}))

	app.Use(func(c *fiber.Ctx) error {
		return c.
			Status(fiber.StatusNotFound).
			Render("views/errors/404", fiber.Map{
				"ApplicationName":       config.GetString("APP_NAME"),
				"ApplicationDescriptor": config.GetString("APP_DESCRIPTOR"),
			}, "views/layouts/guest")
	})

	return app
}

func createNewFiberApp(config ServerConfig) *fiber.App {
	engine := jet.NewFileSystem(http.FS(config.ViewFilesystem), ".jet.html")

	return fiber.New(fiber.Config{
		AppName:      "Custodian",
		ServerHeader: "Custodian Web Server",
		Views:        engine,
	})
}
