package session

import (
	"context"
	"io"
	"log/slog"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
	"github.com/senither/custodian/database"
	"github.com/senither/custodian/database/model"
	"github.com/senither/custodian/database/repository"
	"github.com/stretchr/testify/assert"
)

func setupSessionStorageMiddleware(c *fiber.Ctx) error {
	sess, err := LoadSessionFromContext(c)
	if err != nil {
		return err
	}

	defer sess.Save()
	c.Locals("session", sess)

	return c.Next()
}

func loadSessionDataAndReturnAsJson(c *fiber.Ctx) error {
	sess, _ := LoadSessionFromContext(c)

	data := make(map[string]interface{})
	for _, key := range sess.Keys() {
		data[key] = sess.Get(key)
	}

	return c.JSON(data)
}

func TestInitiateSessionStorage(t *testing.T) {
	database.InitiateDatabaseConnection(database.MemorySQLiteDSN)

	t.Run("it initiates the session storage", func(t *testing.T) {
		InitiateSessionStorage()

		assert.NotNil(t, store)
		assert.IsType(t, &session.Store{}, store)
	})

	database.Disconnect()
}

func TestLoadSessionFromContext(t *testing.T) {
	database.InitiateDatabaseConnection(database.MemorySQLiteDSN)
	InitiateSessionStorage()

	app := fiber.New()

	app.Get("/", loadSessionDataAndReturnAsJson)
	app.Post("/", func(c *fiber.Ctx) error {
		sess, _ := LoadSessionFromContext(c)

		sess.Set("test", "value")
		sess.Save()

		return c.SendStatus(200)
	})

	t.Run("it can load empty session from context", func(t *testing.T) {
		req := httptest.NewRequest("GET", "/", nil)
		resp, err := app.Test(req)

		body, _ := io.ReadAll(resp.Body)

		assert.Nil(t, err)
		assert.Equal(t, 200, resp.StatusCode)
		assert.JSONEq(t, `{}`, string(body))
	})

	t.Run("it can save and get value in session", func(t *testing.T) {
		req := httptest.NewRequest("POST", "/", nil)
		resp, err := app.Test(req)

		assert.Nil(t, err)
		assert.Equal(t, 200, resp.StatusCode)
		assert.Contains(t, resp.Header.Get("Set-Cookie"), "custodian_session=")

		cookie := strings.Split(resp.Header.Get("Set-Cookie"), ";")[0]

		req = httptest.NewRequest("GET", "/", nil)
		req.Header.Set("Cookie", cookie)
		resp, err = app.Test(req)

		body, _ := io.ReadAll(resp.Body)

		assert.Nil(t, err)
		assert.Equal(t, 200, resp.StatusCode)
		assert.JSONEq(t, `{"test":"value"}`, string(body))
	})

	database.Disconnect()
}

func TestGetSessionFromContext(t *testing.T) {
	database.InitiateDatabaseConnection(database.MemorySQLiteDSN)
	InitiateSessionStorage()

	app := fiber.New()

	app.Use(setupSessionStorageMiddleware)
	app.Get("/", func(c *fiber.Ctx) error {
		_, err := GetSessionFromContext(c)
		if err != nil {
			return err
		}

		return c.SendStatus(200)
	})

	t.Run("it can get the session from context", func(t *testing.T) {
		req := httptest.NewRequest("GET", "/", nil)
		resp, err := app.Test(req)

		assert.Nil(t, err)
		assert.Equal(t, 200, resp.StatusCode)
	})

	t.Run("it returns an error if session is not found in context", func(t *testing.T) {
		app.Get("/error", func(c *fiber.Ctx) error {
			c.Locals("session", nil)
			_, err := GetSessionFromContext(c)

			return err
		})

		req := httptest.NewRequest("GET", "/error", nil)
		resp, err := app.Test(req)

		assert.Nil(t, err)
		assert.Equal(t, 401, resp.StatusCode)
	})

	database.Disconnect()
}

func TestSetAuthenticatedUser(t *testing.T) {
	database.InitiateDatabaseConnection(database.MemorySQLiteDSN)
	InitiateSessionStorage()

	app := fiber.New()

	app.Use(setupSessionStorageMiddleware)
	app.Get("/", loadSessionDataAndReturnAsJson)
	app.Post("/set-user", func(c *fiber.Ctx) error {
		user := model.User{}
		user.ID = 1

		err := SetAuthenticatedUser(c, user)
		if err != nil {
			return err
		}

		return c.SendStatus(fiber.StatusOK)
	})

	t.Run("it sets authenticated user in session", func(t *testing.T) {
		req := httptest.NewRequest("POST", "/set-user", nil)
		resp, err := app.Test(req)

		assert.Nil(t, err)
		assert.Equal(t, 200, resp.StatusCode)
		assert.Contains(t, resp.Header.Get("Set-Cookie"), "custodian_session=")

		cookie := strings.Split(resp.Header.Get("Set-Cookie"), ";")[0]

		req = httptest.NewRequest("GET", "/", nil)
		req.Header.Set("Cookie", cookie)
		resp, err = app.Test(req)

		body, _ := io.ReadAll(resp.Body)

		assert.Nil(t, err)
		assert.Equal(t, 200, resp.StatusCode)
		assert.JSONEq(t, `{"_internal.UID":1}`, string(body))
	})

	t.Run("it returns an error if session is not found in context", func(t *testing.T) {
		app.Post("/error", func(c *fiber.Ctx) error {
			c.Locals("session", nil)
			user := model.User{}
			user.ID = 1

			return SetAuthenticatedUser(c, user)
		})

		req := httptest.NewRequest("POST", "/error", nil)
		resp, err := app.Test(req)

		assert.Nil(t, err)
		assert.Equal(t, 401, resp.StatusCode)
	})

	database.Disconnect()
}

func TestGetAuthenticatedUser(t *testing.T) {
	database.InitiateDatabaseConnection(database.MemorySQLiteDSN)
	InitiateSessionStorage()

	err := repository.CreateUserWithoutPasswordEncryption(context.Background(), model.User{
		Name:     "Test User",
		Email:    "test@example.com",
		Password: "password",
	})

	assert.Nil(t, err)

	app := fiber.New()

	app.Use(setupSessionStorageMiddleware)
	app.Post("/set-user", func(c *fiber.Ctx) error {
		user, err := repository.FindUserByEmail(c.UserContext(), "test@example.com")
		if err != nil {
			return err
		}

		err = SetAuthenticatedUser(c, user)
		if err != nil {
			return err
		}

		return c.SendStatus(fiber.StatusOK)
	})

	app.Get("/get-user", func(c *fiber.Ctx) error {
		user, err := GetAuthenticatedUser(c)
		if err != nil {
			return err
		}

		return c.JSON(user)
	})

	t.Run("it gets authenticated user from session", func(t *testing.T) {
		// Set user in session
		req := httptest.NewRequest("POST", "/set-user", nil)
		resp, err := app.Test(req)

		responseBody, _ := io.ReadAll(resp.Body)
		slog.Error("failed test", "body", responseBody, "err", err)

		assert.Nil(t, err)
		assert.Equal(t, 200, resp.StatusCode)

		cookie := strings.Split(resp.Header.Get("Set-Cookie"), ";")[0]

		// Get user from session
		req = httptest.NewRequest("GET", "/get-user", nil)
		req.Header.Set("Cookie", cookie)
		resp, err = app.Test(req)

		body, _ := io.ReadAll(resp.Body)
		assert.Nil(t, err)
		assert.Equal(t, 200, resp.StatusCode)
		assert.Contains(t, string(body), `"ID":1`)
		assert.Contains(t, string(body), `"name":"Test User"`)
		assert.Contains(t, string(body), `"email":"test@example.com"`)
	})

	t.Run("it redirects user to login page if user is not authenticated", func(t *testing.T) {
		req := httptest.NewRequest("GET", "/get-user", nil)
		resp, err := app.Test(req)

		assert.Nil(t, err)
		assert.Equal(t, 302, resp.StatusCode)
		assert.Equal(t, "/login", resp.Header.Get("Location"))
	})
}

func TestGetAuthenticatedUserWithoutRedirects(t *testing.T) {
	database.InitiateDatabaseConnection(database.MemorySQLiteDSN)
	InitiateSessionStorage()

	err := repository.CreateUserWithoutPasswordEncryption(context.Background(), model.User{
		Name:     "Test User",
		Email:    "test@example.com",
		Password: "password",
	})

	assert.Nil(t, err)

	app := fiber.New()

	app.Use(setupSessionStorageMiddleware)
	app.Post("/set-user", func(c *fiber.Ctx) error {
		user, err := repository.FindUserByEmail(c.UserContext(), "test@example.com")
		if err != nil {
			return err
		}

		err = SetAuthenticatedUser(c, user)
		if err != nil {
			return err
		}

		return c.SendStatus(fiber.StatusOK)
	})

	app.Get("/get-user", func(c *fiber.Ctx) error {
		user, err := GetAuthenticatedUserWithoutRedirects(c)
		if err != nil {
			return err
		}

		return c.JSON(user)
	})

	t.Run("it gets authenticated user from session without redirects", func(t *testing.T) {
		// Set user in session
		req := httptest.NewRequest("POST", "/set-user", nil)
		resp, err := app.Test(req)

		responseBody, _ := io.ReadAll(resp.Body)
		slog.Error("failed test", "body", responseBody, "err", err)

		assert.Nil(t, err)
		assert.Equal(t, 200, resp.StatusCode)

		cookie := strings.Split(resp.Header.Get("Set-Cookie"), ";")[0]

		// Get user from session
		req = httptest.NewRequest("GET", "/get-user", nil)
		req.Header.Set("Cookie", cookie)
		resp, err = app.Test(req)

		body, _ := io.ReadAll(resp.Body)
		assert.Nil(t, err)
		assert.Equal(t, 200, resp.StatusCode)
		assert.Contains(t, string(body), `"ID":1`)
	})

	t.Run("it returns error if user is not authenticated", func(t *testing.T) {
		req := httptest.NewRequest("GET", "/get-user", nil)
		resp, err := app.Test(req)

		assert.Nil(t, err)
		assert.Equal(t, 500, resp.StatusCode)
	})

	database.Disconnect()
}

func TestGetSessionStore(t *testing.T) {
	database.InitiateDatabaseConnection(database.MemorySQLiteDSN)
	InitiateSessionStorage()

	t.Run("it returns the session store instance", func(t *testing.T) {
		assert.NotNil(t, GetSessionStore())
		assert.IsType(t, &session.Store{}, GetSessionStore())
	})

	database.Disconnect()
}
