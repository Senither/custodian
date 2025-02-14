package templating

import (
	"embed"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/jet/v2"
	"github.com/senither/custodian/config"
	"github.com/stretchr/testify/assert"
)

func TestNewTemplatingEngine(t *testing.T) {
	cfg := config.ServerConfig{
		ViewFilesystem: embed.FS{},
	}

	t.Run("it returns a new templating engine", func(t *testing.T) {
		engine := NewTemplatingEngine(cfg)

		assert.NotNil(t, engine)
		assert.IsType(t, &jet.Engine{}, engine)
	})
}

func TestRegisterTemplateDefaults(t *testing.T) {
	cfg := config.ServerConfig{
		ViewFilesystem: embed.FS{},
	}

	engine := NewTemplatingEngine(cfg)

	fiber.New(fiber.Config{
		Views: engine,
	})

	t.Run("it registers the default template bindings", func(t *testing.T) {
		RegisterTemplateDefaults(engine)

		_, foundErrors := engine.Templates.LookupGlobal("errors")
		assert.True(t, foundErrors, "'errors' global not found")

		_, foundOld := engine.Templates.LookupGlobal("old")
		assert.True(t, foundOld, "'old' global not found")
	})
}
