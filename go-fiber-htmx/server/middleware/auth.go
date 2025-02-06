package middleware

import (
	"github.com/gofiber/fiber/v2"
	"github.com/senither/custodian/server/session"
	"github.com/senither/custodian/server/utils"
)

func Authenticated() fiber.Handler {
	return func(c *fiber.Ctx) error {
		user, _ := session.GetAuthenticatedUserWithoutRedirects(c)
		if user == nil {
			return utils.RedirectWithHtmx(c, "/login")
		}

		return c.Next()
	}
}

func Guest() fiber.Handler {
	return func(c *fiber.Ctx) error {
		user, _ := session.GetAuthenticatedUserWithoutRedirects(c)
		if user != nil {
			return utils.RedirectWithHtmx(c, "/dashboard")
		}

		return c.Next()
	}
}
