package config

import "embed"

type ServerConfig struct {
	PublicFilesystem embed.FS
	ViewFilesystem   embed.FS
}
