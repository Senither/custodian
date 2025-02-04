package server

import (
	"embed"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/filesystem"
	"github.com/gofiber/template/jet/v2"
	"github.com/senither/custodian/server/router"
)

type ServerConfig struct {
	PublicFilesystem embed.FS
	ViewFilesystem   embed.FS
}

func NewServer(config ServerConfig) *fiber.App {
	app := createNewFiberApp(config)

	router.RegisterRoutes(app)

	app.Use("", filesystem.New(filesystem.Config{
		Root: http.FS(config.PublicFilesystem),
	}))

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
