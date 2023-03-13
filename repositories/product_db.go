package repositories

import (
	"gorm.io/gorm"
)

type productReposityDB struct {
	db *gorm.DB
}

func NewProductRepositoryDB(db *gorm.DB) ProductReposity {
	db.AutoMigrate(&product{});
	mockData(db);
	return productReposityDB{db:db}; 
}



func (r productReposityDB) GetProduct() ([]product, error) {
	
	var product []product;

	err := r.db.Order("quantity desc").Limit(30).Find(&product).Error;
	
	return product, err;
}