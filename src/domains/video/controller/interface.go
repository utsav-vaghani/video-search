package controller

import "github.com/gofiber/fiber/v2"

// Video defining the methods available for controller layer
type Video interface {
	Get(ctx *fiber.Ctx) error
}
