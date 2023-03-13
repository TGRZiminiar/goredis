package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"goredis/services"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/gofiber/fiber/v2"
)

type catalogHandlerRedis struct {
	catalogSrv services.CatalogService;
	redisClient	*redis.Client;
}

func NewCatalogHandlerRedis(catalogSrv services.CatalogService, redisClient *redis.Client) CatalogHandler {
	return catalogHandlerRedis{catalogSrv, redisClient};
}

func (h catalogHandlerRedis) GetProduct(c *fiber.Ctx) error {
	
	key := "handler::GetProduct";

	//Redis Get
	if responseJson, err := h.redisClient.Get(context.Background(), key).Result(); err == nil {
		fmt.Println("redis")
		c.Set("Content-Type", "application/json")
		return c.SendString(responseJson)
	}


	// Service
	products, err := h.catalogSrv.GetProduct()
	if err != nil {
		return err;
	}

	response := fiber.Map{
		"status":"ok",
		"products":products,
	}

	//Redis Set ไม่มีปัญหาให้เซ็ตลง
	if data, err := json.Marshal(response); err == nil {
		h.redisClient.Set(context.Background(), key, string(data), time.Second * 10);
	}

	fmt.Println("database");
	return c.JSON(response);

}
