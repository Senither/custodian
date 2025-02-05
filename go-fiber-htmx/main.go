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
	cfgErr := config.LoadConfig()
	if cfgErr != nil {
		panic(cfgErr)
	}

	dbErr := database.InitiateDatabaseConnection()
	if dbErr != nil {
		panic(dbErr)
	}

	app := server.NewServer(server.ServerConfig{
		ViewFilesystem:   views,
		PublicFilesystem: public,
	})

	app.Listen(config.Get().Application.Address)
}
