package middleware

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/senither/custodian/server/session"
)

var largestUserLogSizeFound = 0

func newFiberLogger() fiber.Handler {
	return logger.New(logger.Config{
		Format: "${time} | ${status} | ${latency} | ${ip} | ${user} | ${method} | ${path} | ${error} \n",
		CustomTags: map[string]logger.LogFunc{
			"user": handleLoggingAuthenticatedUser,
		},
	})
}

func handleLoggingAuthenticatedUser(output logger.Buffer, c *fiber.Ctx, data *logger.Data, extraParam string) (int, error) {
	userStr := "guest"

	user, err := session.GetAuthenticatedUserWithoutRedirects(c)
	if err == nil {
		userStr = "user:" + strconv.FormatUint(uint64(user.ID), 10)
	}

	if len(userStr) > largestUserLogSizeFound {
		largestUserLogSizeFound = len(userStr)
	}

	return output.WriteString(maskString(userStr, largestUserLogSizeFound))
}

func maskString(str string, size int) string {
	if len(str) >= size {
		return str
	}

	return maskString(" "+str, size)
}
