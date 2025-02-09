package templating

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/jet/v2"
	"github.com/senither/custodian/config"
)

func NewTemplatingEngine(cfg config.ServerConfig) *jet.Engine {
	engine := jet.NewFileSystem(http.FS(cfg.ViewFilesystem), ".jet.html")

	engine.AddFunc("getInputErrorFor", getInputErrorFor)
	engine.AddFunc("getOldInputValueFor", getOldInputValueFor)

	return engine
}

// Registers the default template bindings that are expected to exist in the
// templates, but can be overridden by the route handler if needed.
//
// Note: This should be called after the application has been instantiated
// so the templates can be loaded and bound to the engine.
func RegisterTemplateDefaults(engine *jet.Engine) {
	engine.Templates.AddGlobal("errors", &fiber.Map{})
	engine.Templates.AddGlobal("old", &fiber.Map{})
}
