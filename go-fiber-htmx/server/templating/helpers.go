package templating

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func getInputErrorFor(errors *fiber.Map, key string) string {
	if errors == nil {
		return ""
	}

	errs, ok := (*errors)[key].([]string)
	if !ok || len(errs) == 0 {
		return ""
	}

	lines := ""
	for _, err := range errs {
		lines += "<li>" + err + "</li>"
	}

	return fmt.Sprintf(
		"<div class=\"mt-2\"><ul class=\"text-error text-sm\">%s</ul></div>",
		lines,
	)
}
