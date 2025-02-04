package main

import (
	"embed"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/filesystem"
	"github.com/gofiber/template/jet/v2"
	"github.com/senither/custodian/config"
)

//go:embed views/**
var views embed.FS

//go:embed public/**
var public embed.FS

func main() {
	engine := jet.NewFileSystem(http.FS(views), ".jet.html")

	app := fiber.New(fiber.Config{
		AppName:      "Custodian",
		ServerHeader: "Custodian Web Server",
		Views:        engine,
	})

	app.Get("/", func(c *fiber.Ctx) error {
		return c.Render("views/index", fiber.Map{}, "views/layouts/app")
	})

	app.Get("/login", func(c *fiber.Ctx) error {
		return c.Render("views/login", fiber.Map{}, "views/layouts/guest")
	})

	app.Use("", filesystem.New(filesystem.Config{
		Root: http.FS(public),
	}))

	app.Listen(config.GetString("APP_ADDR"))
}
