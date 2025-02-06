package session

import (
	"context"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
	"github.com/senither/custodian/database"
	"github.com/senither/custodian/database/model"
	"github.com/senither/custodian/database/repository"
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

// Sets the authenticated user in the session store using the provided user model.
// This function should only be called within the request cycle, so that the session middleware saves the session to the database.
func SetAuthenticatedUser(ctx *fiber.Ctx, user model.User) error {
	session, err := GetSessionFromContext(ctx)
	if err != nil {
		return err
	}

	session.Set("_internal.UID", user.ID)

	return nil
}

// Gets the authenticated user from the session store using the session ID stored in the cookie from the provided context.
// This function returns an error if the user is not found in the session store, or if the user ID stored in the session is not a valid user ID.
func GetAuthenticatedUser(ctx *fiber.Ctx) (*model.User, error) {
	value := ctx.Locals("user")

	localUser, ok := value.(*model.User)
	if ok {
		return localUser, nil
	}

	session, err := GetSessionFromContext(ctx)
	if err != nil {
		ctx.Append("HX-Redirect", "/login")
		return nil, err
	}

	uid, ok := session.Get("_internal.UID").(uint)
	if !ok {
		ctx.Append("HX-Redirect", "/login")
		return nil, fiber.ErrUnauthorized
	}

	dbUser, err := repository.FindUserByID(ctx.UserContext(), uid)
	if err != nil {
		ctx.Append("HX-Redirect", "/login")
		return nil, fiber.ErrUnauthorized
	}

	ctx.Locals("user", dbUser)

	return &dbUser, nil
}

func GetSessionStore() *session.Store {
	return store
}
