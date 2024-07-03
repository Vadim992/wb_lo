package server

import (
	"github.com/gofiber/fiber"
)

func NewServer() *fiber.App {
	return fiber.New()
}
