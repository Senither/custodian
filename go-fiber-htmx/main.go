package main

import (
	"embed"

	"github.com/senither/custodian/config"
	"github.com/senither/custodian/server"
)

//go:embed views/**
var views embed.FS

//go:embed public/**
var public embed.FS

func main() {
	app := server.NewServer(server.ServerConfig{
		ViewFilesystem:   views,
		PublicFilesystem: public,
	})

	app.Listen(config.GetString("APP_ADDR"))
}
