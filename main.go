package main

import (
	"goredis/handlers"
	"goredis/repositories"
	"goredis/services"

	"github.com/go-redis/redis/v8"
	"github.com/gofiber/fiber/v2"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {

	
	// productRepo := repositories.NewProductRepositoryDB()
	
	db := initDatabase();
	redisClient := initRedis();
	_ = redisClient;

	// productRepo := repositories.NewProductRepositoryDB(db);
	productRepo := repositories.NewProductRepositoryRedis(db, redisClient);
	
	productService := services.NewCatalogService(productRepo);
	// productHandler := handlers.NewCatalogHandlerRedis(productService, redisClient);
	productHandler := handlers.NewCatalogHandler(productService);

	
	app := fiber.New();

	app.Get("/products", productHandler.GetProduct);

	app.Listen(":5000");

}

func initDatabase() *gorm.DB {

	dial := mysql.Open("root:123@tcp(localhost:3306)/testRedis")
	db, err := gorm.Open(dial, &gorm.Config{})
	if err != nil {
		panic(err)
	}
	return db

}	

func initRedis() *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})
}

// docker compose run --rm k6 run /scripts/test.js docker run --rm -v /scripts:/scripts loadimpact/k6 run /scripts/test.js

