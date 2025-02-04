package main

import (
	"embed"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/filesystem"
	"github.com/gofiber/template/html/v2"
	"github.com/senither/custodian/config"
)

//go:embed views/**
var views embed.FS

//go:embed public/**
var public embed.FS

func main() {
	engine := html.NewFileSystem(http.FS(views), ".html")

	app := fiber.New(fiber.Config{
		AppName:      "Custodian",
		ServerHeader: "Custodian Web Server",
		Views:        engine,
	})

	app.Get("/", func(c *fiber.Ctx) error {
		return c.Render("views/index", fiber.Map{})
	})

	app.Use("", filesystem.New(filesystem.Config{
		Root: http.FS(public),
	}))

	app.Listen(config.GetString("APP_ADDR"))
}
