package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/senither/custodian/database/repository"
	"github.com/senither/custodian/server/security"
	"github.com/senither/custodian/server/session"
	"github.com/senither/custodian/server/utils"
	"github.com/senither/custodian/server/validator"
)

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
