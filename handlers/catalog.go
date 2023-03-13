package handlers

import "github.com/gofiber/fiber/v2"

type CatalogHandler interface {
	GetProduct(c *fiber.Ctx) error
}

