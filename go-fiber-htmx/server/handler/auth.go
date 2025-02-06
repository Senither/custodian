package handler

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/senither/custodian/database/model"
	"github.com/senither/custodian/database/repository"
	"github.com/senither/custodian/server/security"
	"github.com/senither/custodian/server/session"
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

	user, err := repository.FindUserByEmail(c.UserContext(), loginRequest.Email)
	if err != nil {
		return c.Render("views/auth/login", fiber.Map{
			"errorMessage": "Invalid email or password (1)",
		})
	}

	if !security.VerifyPassword(user.Password, loginRequest.Password) {
		return c.Render("views/auth/login", fiber.Map{
			"errorMessage": "Invalid email or password (2)",
		})
	}

	session.SetAuthenticatedUser(c, user)

	c.Append("HX-Redirect", "/dashboard")
	return c.SendString("Login successful")
}

type RegisterRequest struct {
	Name            string `validate:"required,min=3,max=80"`
	Email           string `validate:"required,email"`
	Password        string `validate:"required,min=8"`
	PasswordConfirm string `validate:"required"`
}

func Register(c *fiber.Ctx) error {
	registerRequest := RegisterRequest{
		Name:            c.FormValue("name"),
		Email:           c.FormValue("email"),
		Password:        c.FormValue("password"),
		PasswordConfirm: c.FormValue("password_confirm"),
	}

	if err := validator.Parse(c.UserContext(), registerRequest); err != nil {
		return c.Render("views/auth/register", fiber.Map{
			"errors": err,
		})
	}

	if registerRequest.Password != registerRequest.PasswordConfirm {
		return c.Render("views/auth/register", fiber.Map{
			"errors": fiber.Map{
				"password": []string{"Passwords do not match"},
			},
		})
	}

	if repository.UserExistsByEmail(c.UserContext(), registerRequest.Email) {
		return c.Render("views/auth/register", fiber.Map{
			"errors": fiber.Map{
				"email": []string{"Email is already in use"},
			},
		})
	}

	currentTime := time.Now()
	createErr := repository.CreateUser(c.UserContext(), model.User{
		Name:            registerRequest.Name,
		Email:           registerRequest.Email,
		EmailVerifiedAt: &currentTime,
		Password:        registerRequest.Password,
	})

	if createErr != nil {
		return c.Render("views/auth/register", fiber.Map{
			"errorMessage": "Failed to create user",
		})
	}

	user, _ := repository.FindUserByEmail(c.UserContext(), registerRequest.Email)
	session.SetAuthenticatedUser(c, user)

	c.Append("HX-Redirect", "/dashboard")
	return c.SendString("Registration successful")
}

func ForgotPassword(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"page": "Forgot Password",
	})
}
