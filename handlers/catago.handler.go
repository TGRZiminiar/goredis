package handlers

import (
	"goredis/services"

	"github.com/gofiber/fiber/v2"
)

type catalogHandler struct {
	catalogSrv services.CatalogService
}


func NewCatalogHandler(catalogSrv services.CatalogService) CatalogHandler {
	return catalogHandler{catalogSrv};
}

func (h catalogHandler) GetProduct(c *fiber.Ctx) error {

	products, err := h.catalogSrv.GetProduct();
	if err != nil {
		return err;
	}

	response := fiber.Map{
		"status":"ok",
		"products":products,
	}

	return c.JSON(response);
} 
