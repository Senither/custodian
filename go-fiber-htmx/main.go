package main

import (
	"embed"
	"log/slog"

	"github.com/senither/custodian/config"
	"github.com/senither/custodian/database"
	"github.com/senither/custodian/server"
)

//go:embed views/**
var views embed.FS

//go:embed public/**
var public embed.FS

func main() {
	cfgErr := config.LoadConfig(".env")
	if cfgErr != nil {
		panic(cfgErr)
	}

	if config.Get().Application.Debug {
		slog.Info("Starting application with debug mode enabled")
	}

	dbErr := database.InitiateDatabaseConnection(config.Get().Database.Url)
	if dbErr != nil {
		panic(dbErr)
	}

	app := server.NewServer(config.ServerConfig{
		ViewFilesystem:   views,
		PublicFilesystem: public,
	})

	app.Listen(config.Get().Application.Address)
}
