package web

import "github.com/gofiber/fiber/v2"

func Api(c *fiber.Ctx) error {
	return c.Next()
}

func Container(c *fiber.Ctx) error {
	return c.Next()
}

func Network(c *fiber.Ctx) error {
	return c.Next()
}
