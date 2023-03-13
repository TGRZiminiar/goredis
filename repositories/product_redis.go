package repositories

import (
	"context"
	"encoding/json"
	"time"

	"github.com/go-redis/redis/v8"
	"gorm.io/gorm"
)

type productReposityRedis struct {
	db *gorm.DB;
	redisClient *redis.Client
}

func NewProductRepositoryRedis(db *gorm.DB, redisClient *redis.Client) ProductReposity  {
	db.AutoMigrate(&product{});
	mockData(db);
	return productReposityRedis{db, redisClient};
}

func (r productReposityRedis) GetProduct() ([]product, error) {
	
	key := "repository::GetProduct";
	var product []product;

	//Redis Get
	productJson, err := r.redisClient.Get(context.Background(), key).Result();
	//ถ้าไม่มี error แปลว่ามีค่า
	if err == nil {
		err = json.Unmarshal([]byte(productJson), &product);
		if err == nil {
			// fmt.Println("Redis");
			return product, nil;
		}
	}

	//database
	err = r.db.Order("quantity desc").Limit(30).Find(&product).Error;
	if err != nil {
		return nil, err;
	}

	//Redis Set ต้องแปลงเป็น JSON ก่อน
	data, err := json.Marshal(product);
	if err != nil {
		return nil, err;
	}

	err = r.redisClient.Set(context.Background(), key, string(data), time.Second * 10).Err();
	if err != nil {
		return nil, err;
	}

	// fmt.Println("database")

	return product,err;
}



