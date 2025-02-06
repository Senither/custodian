package session

import (
	"context"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
	"github.com/senither/custodian/database"
	"github.com/senither/custodian/database/model"
)

var store *session.Store

func InitiateSessionStorage() {
	config := session.Config{
		Expiration: 24 * time.Hour,
		KeyLookup:  "cookie:custodian_session",
	}

	config.Storage = model.NewDatabaseSessionStorage(
		database.GetConnectionWithContext(context.Background()),
		config,
	)

	store = session.New(config)
}

// Loads the users session from the database using the session ID stored in the cookie from the provided context.
func LoadSessionFromContext(ctx *fiber.Ctx) (*session.Session, error) {
	return store.Get(ctx)
}

// Gets the current session from the context locals, and returns an error if the session is not found.
// This function is used to get the session from the context after it has been loaded by the LoadSessionFromContext function.
func GetSessionFromContext(ctx *fiber.Ctx) (*session.Session, error) {
	value := ctx.Locals("session")

	session, ok := value.(*session.Session)
	if !ok {
		return nil, fiber.ErrUnauthorized
	}

	return session, nil
}

func GetSessionStore() *session.Store {
	return store
}
