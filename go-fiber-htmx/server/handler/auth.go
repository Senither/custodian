package handler

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/senither/custodian/database/model"
	"github.com/senither/custodian/database/repository"
	"github.com/senither/custodian/server/security"
	"github.com/senither/custodian/server/session"
	"github.com/senither/custodian/server/utils"
	"github.com/senither/custodian/server/validator"
)

type LoginRequest struct {
	Email    string `validate:"required,email"`
	Password string `validate:"required"`
	Remember bool   `validate:"boolean"`
}

func Login(c *fiber.Ctx) error {
	request := LoginRequest{
		Email:    c.FormValue("email"),
		Password: c.FormValue("password"),
		Remember: c.FormValue("remember") == "on",
	}

	if err := validator.Parse(c.UserContext(), request); err != nil {
		return c.Render("views/auth/login", fiber.Map{
			"errors": err,
		})
	}

	user, err := repository.FindUserByEmail(c.UserContext(), request.Email)
	if err != nil {
		return c.Render("views/auth/login", fiber.Map{
			"errorMessage": "Invalid email or password",
		})
	}

	if !security.VerifyPassword(user.Password, request.Password) {
		return c.Render("views/auth/login", fiber.Map{
			"errorMessage": "Invalid email or password",
		})
	}

	session.SetAuthenticatedUser(c, user)

	return utils.RedirectWithHtmx(c, "/dashboard")
}

func Logout(c *fiber.Ctx) error {
	ses, err := session.GetSessionFromContext(c)
	if err != nil {
		return utils.RedirectWithHtmx(c, "/login")
	}

	ses.Destroy()

	return utils.RedirectWithHtmx(c, "/login")
}

type RegisterRequest struct {
	Name            string `validate:"required,min=3,max=80"`
	Email           string `validate:"required,email"`
	Password        string `validate:"required,min=8"`
	PasswordConfirm string `validate:"required"`
}

func Register(c *fiber.Ctx) error {
	request := RegisterRequest{
		Name:            c.FormValue("name"),
		Email:           c.FormValue("email"),
		Password:        c.FormValue("password"),
		PasswordConfirm: c.FormValue("password_confirm"),
	}

	if err := validator.Parse(c.UserContext(), request); err != nil {
		return c.Render("views/auth/register", fiber.Map{
			"old":    utils.ConvertToFiberMap(request),
			"errors": err,
		})
	}

	if request.Password != request.PasswordConfirm {
		return c.Render("views/auth/register", fiber.Map{
			"old": utils.ConvertToFiberMap(request),
			"errors": &fiber.Map{
				"password": []string{"Passwords do not match"},
			},
		})
	}

	if repository.UserExistsByEmail(c.UserContext(), request.Email) {
		return c.Render("views/auth/register", fiber.Map{
			"old": utils.ConvertToFiberMap(request),
			"errors": &fiber.Map{
				"email": []string{"Email is already in use"},
			},
		})
	}

	currentTime := time.Now()
	createErr := repository.CreateUser(c.UserContext(), model.User{
		Name:            request.Name,
		Email:           request.Email,
		EmailVerifiedAt: &currentTime,
		Password:        request.Password,
	})

	if createErr != nil {
		return c.Render("views/auth/register", fiber.Map{
			"old":          utils.ConvertToFiberMap(request),
			"errorMessage": "Failed to create user",
		})
	}

	user, _ := repository.FindUserByEmail(c.UserContext(), request.Email)
	session.SetAuthenticatedUser(c, user)

	return utils.RedirectWithHtmx(c, "/dashboard")
}

type ForgotPasswordRequest struct {
	Email string `validate:"required,email"`
}

func ForgotPassword(c *fiber.Ctx) error {
	request := ForgotPasswordRequest{
		Email: c.FormValue("email"),
	}

	if err := validator.Parse(c.UserContext(), request); err != nil {
		return c.Render("views/auth/forgot-password", fiber.Map{
			"errors": err,
		})
	}

	if !repository.UserExistsByEmail(c.UserContext(), request.Email) {
		return c.Render("views/auth/forgot-password", fiber.Map{
			"errors": &fiber.Map{
				"email": []string{"No user with that email exists"},
			},
		})
	}

	// Here is where we'd send the email to the user with the password reset link,
	// this could be done with any email service that allows sending emails
	// via a HTTP API (such as Resend, SendGrid, Mailgun, etc).
	// However, for this example we will just return a success message.

	return c.Render("views/auth/forgot-password", fiber.Map{
		"actionMessage": "A password reset link has been sent to your email",
	})
}
