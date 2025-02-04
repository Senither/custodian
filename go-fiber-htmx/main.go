package main

import (
	"embed"

	"github.com/senither/custodian/config"
	"github.com/senither/custodian/database"
	"github.com/senither/custodian/server"
)

//go:embed views/**
var views embed.FS

//go:embed public/**
var public embed.FS

func main() {
	err := database.InitiateDatabaseConnection()
	if err != nil {
		panic(err)
	}

	app := server.NewServer(server.ServerConfig{
		ViewFilesystem:   views,
		PublicFilesystem: public,
	})

	app.Listen(config.GetString("APP_ADDR"))
}
