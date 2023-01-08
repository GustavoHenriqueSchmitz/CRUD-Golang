package models

import "github.com/gofiber/fiber/v2"

type Server struct {
	App  *fiber.App
	Port string
}
