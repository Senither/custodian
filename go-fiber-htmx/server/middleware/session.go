package middleware

import (
	"math/rand"
	"strings"

	"github.com/gofiber/fiber/v2"
	fiberSession "github.com/gofiber/fiber/v2/middleware/session"
	"github.com/senither/custodian/server/session"
)

func handleSessions(c *fiber.Ctx) error {
	ses, err := session.LoadSessionFromContext(c)

	if err != nil {
		if strings.Contains(err.Error(), "failed to decode session data") {
			c.ClearCookie("custodian_session")
			return c.Next()
		}

		return err
	}

	c.Locals("session", ses)

	originalValues := getSessionDataAsMap(ses)
	defer saveSessionOnChanges(ses, originalValues)

	return c.Next()
}

func saveSessionOnChanges(ses *fiberSession.Session, original map[string]interface{}) {
	// Gives a 0.5% chance to save the session on each request regardless of if the session
	// data has changed or not, this is to help extend the session "expires at" time even
	// if the session is not modified within the expiration period.
	if rand.Intn(199) == 0 {
		ses.Save()
		return
	}

	// If the amount of session keys has changed we know that something has
	// been added or removed within the session, so we can save early.
	if len(ses.Keys()) != len(original) {
		ses.Save()
		return
	}

	// Checks if the session data has changed between the start and end of
	// the request, if the session data was changed we save the session.
	for key, value := range getSessionDataAsMap(ses) {
		val, ok := original[key]
		if !ok || val != value {
			ses.Save()
			return
		}
	}
}

func getSessionDataAsMap(ses *fiberSession.Session) map[string]interface{} {
	data := make(map[string]interface{}, len(ses.Keys()))

	for _, key := range ses.Keys() {
		data[key] = ses.Get(key)
	}

	return data
}
