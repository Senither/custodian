package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/senither/custodian/server/validator"
)

type LoginRequest struct {
	Email    string `validate:"required,email"`
	Password string `validate:"required"`
	Remember bool   `validate:"boolean"`
}

func Login(c *fiber.Ctx) error {
	loginRequest := LoginRequest{
		Email:    c.FormValue("email"),
		Password: c.FormValue("password"),
		Remember: c.FormValue("remember") == "on",
	}

	if err := validator.Parse(c.UserContext(), loginRequest); err != nil {
		return c.Render("views/auth/login", fiber.Map{
			"errors": err,
		})
	}

	// TODO: Authenticate the user, and redirect them to the dashboard.
	// or fail the login attempt and show an error message to the user.

	// Success state, redirects to the dashboard.
	c.Append("HX-Redirect", "/dashboard")

	// Fail state, shows an error message to the user.
	return c.Render("views/auth/login", fiber.Map{
		"errorMessage": "Invalid email or password",
	})
}

func Register(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"page": "Sign up",
	})
}

func ForgotPassword(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"page": "Forgot Password",
	})
}
