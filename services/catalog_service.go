package services

import "goredis/repositories"

type catalogService struct {
	productRepo repositories.ProductReposity
}

func NewCatalogService(productRepo repositories.ProductReposity) (CatalogService) {
	return catalogService{productRepo};
}

func (s catalogService) GetProduct() ([]Product, error){
	
	productDb, err := s.productRepo.GetProduct();
	if err != nil {
		return nil, err;
	}

	var products []Product;

	for _, p := range productDb {
		products = append(products, Product{
			ID: p.ID,
			Name: p.Name,
			Quantity: p.Quantity,
		})
	}


	return products, nil;
}
