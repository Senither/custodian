package utils

import "github.com/gofiber/fiber/v2"

func GetRequestHeader(c *fiber.Ctx, key string) []string {
	val, ok := c.GetReqHeaders()[key]
	if !ok {
		return make([]string, 0)
	}

	return val
}
