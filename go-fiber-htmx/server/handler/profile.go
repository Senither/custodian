package handler

import (
	"math/rand"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/senither/custodian/database/model"
	"github.com/senither/custodian/database/repository"
	"github.com/senither/custodian/server/security"
	"github.com/senither/custodian/server/session"
	"github.com/senither/custodian/server/utils"
	"github.com/senither/custodian/server/validator"
)

type UpdateProfileInformationRequest struct {
	Name  string `validate:"required,min=3,max=80"`
	Email string `validate:"required,email"`
}

func UpdateProfileInformation(c *fiber.Ctx) error {
	user, err := session.GetAuthenticatedUser(c)
	if err != nil {
		return c.SendString("Failed to get the authenticated user")
	}

	request := UpdateProfileInformationRequest{
		Name:  c.FormValue("name"),
		Email: c.FormValue("email"),
	}

	if err := validator.Parse(c.UserContext(), request); err != nil {
		return c.Render("views/profile", fiber.Map{
			"AuthenticatedUser": user,
			"errors":            err,
		})
	}

	emailUser, emailErr := repository.FindUserByEmail(c.UserContext(), request.Email)
	if emailErr != nil || emailUser.ID != user.ID {
		return c.Render("views/profile", fiber.Map{
			"AuthenticatedUser": user,
			"errors": &fiber.Map{
				"email": []string{"The email address is already in use"},
			},
		})
	}

	dbErr := repository.UpdateUser(c.UserContext(), *user, model.User{
		Name:  request.Name,
		Email: request.Email,
	})

	if dbErr != nil {
		return c.Render("views/profile", fiber.Map{
			"AuthenticatedUser":        user,
			"ActionInformationMessage": "Failed to save the changes",
			"RandomId":                 "update-profile-" + strconv.FormatInt(rand.Int63(), 10),
		})
	}

	user.Name = request.Name
	user.Email = request.Email

	return c.Render("views/profile", fiber.Map{
		"AuthenticatedUser":        user,
		"ActionInformationMessage": "Saved",
		"RandomId":                 "update-profile-" + strconv.FormatInt(rand.Int63(), 10),
	})
}

type UpdateProfilePasswordRequest struct {
	CurrentPassword string `validate:"required,min=1"`
	Password        string `validate:"required,min=8"`
	PasswordConfirm string `validate:"required"`
}

func UpdateProfilePassword(c *fiber.Ctx) error {
	user, err := session.GetAuthenticatedUser(c)
	if err != nil {
		return c.SendString("Failed to get the authenticated user")
	}

	request := UpdateProfilePasswordRequest{
		CurrentPassword: c.FormValue("current_password"),
		Password:        c.FormValue("password"),
		PasswordConfirm: c.FormValue("password_confirm"),
	}

	if err := validator.Parse(c.UserContext(), request); err != nil {
		return c.Render("views/profile", fiber.Map{
			"AuthenticatedUser": user,
			"errors":            err,
		})
	}

	if request.Password != request.PasswordConfirm {
		return c.Render("views/profile", fiber.Map{
			"AuthenticatedUser": user,
			"errors": &fiber.Map{
				"password": []string{"Passwords do not match"},
			},
		})
	}

	if !security.VerifyPassword(user.Password, request.CurrentPassword) {
		return c.Render("views/profile", fiber.Map{
			"AuthenticatedUser": user,
			"errors": &fiber.Map{
				"current_password": []string{"The current password you entered is incorrect"},
			},
		})
	}

	dbErr := repository.UpdateUser(c.UserContext(), *user, model.User{
		Password: request.Password,
	})

	if dbErr != nil {
		return c.Render("views/profile", fiber.Map{
			"AuthenticatedUser":     user,
			"ActionPasswordMessage": "Failed to save the changes",
			"RandomId":              "update-password-" + strconv.FormatInt(rand.Int63(), 10),
		})
	}

	return c.Render("views/profile", fiber.Map{
		"AuthenticatedUser":     user,
		"ActionPasswordMessage": "Saved",
		"RandomId":              "update-password-" + strconv.FormatInt(rand.Int63(), 10),
	})
}

func RenderDeleteAccountModalComponent(c *fiber.Ctx) error {
	return c.Render("views/components/delete-account-modal", fiber.Map{})
}

type DeleteAccountRequest struct {
	Password string `validate:"required"`
}

func DeleteAccount(c *fiber.Ctx) error {
	request := DeleteAccountRequest{
		Password: c.FormValue("password"),
	}

	if err := validator.Parse(c.UserContext(), request); err != nil {
		return c.Render("views/components/delete-account-modal", fiber.Map{
			"errors": err,
		})
	}

	user, err := session.GetAuthenticatedUser(c)
	if err != nil {
		return c.Render("views/components/delete-account-modal", fiber.Map{
			"errors": &fiber.Map{
				"password": []string{"Failed when trying to validate the password"},
			},
		})
	}

	if !security.VerifyPassword(user.Password, request.Password) {
		return c.Render("views/components/delete-account-modal", fiber.Map{
			"errors": &fiber.Map{
				"password": []string{"The password you entered is incorrect"},
			},
		})
	}

	ses, _ := session.GetSessionFromContext(c)
	if ses != nil {
		ses.Destroy()
	}

	go repository.DeleteUserAndRelatedRecordsById(c.UserContext(), user.ID)

	utils.RedirectWithHtmx(c, "/logout")

	return c.Render("views/components/delete-account-modal", fiber.Map{})
}
