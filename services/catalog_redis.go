package services

import (
	"context"
	"encoding/json"
	"fmt"
	"goredis/repositories"
	"time"

	"github.com/go-redis/redis/v8"
)

type catalogServiceRedis struct {
	productRepo repositories.ProductReposity;
	redisClient	*redis.Client;
}

func NewCatalogServiceRedis(productRepo repositories.ProductReposity, r *redis.Client) (catalogServiceRedis) {
	return catalogServiceRedis{
		productRepo: productRepo,
		redisClient: r,
	}
}

func (s *catalogServiceRedis) GetProduct() ([]Product, error) {

	key := "service::GetProduct"
	var products []Product;
	//Redis
	if productJson, err := s.redisClient.Get(context.Background(), key).Result(); err == nil {
		//อ่านได้
		if json.Unmarshal([]byte(productJson), &products) == nil {
			fmt.Println("redis");
			return products,nil;
		}
	}

	//If not exist in redis get data from repo
	productsDb, err := s.productRepo.GetProduct();
	if err != nil {
		return nil, err;
	}

	for _, p := range productsDb {
		products = append(products, Product{
			ID: p.ID,
			Name: p.Name,
			Quantity: p.Quantity,
		})
	}

	//Store In Redis
	if data,err := json.Marshal(products); err == nil {
		s.redisClient.Set(context.Background(), key, string(data), time.Second*10);
	}

	return products, nil;
}