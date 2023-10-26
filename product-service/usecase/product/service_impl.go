package product

import (
	"log"

	"github.com/christiandwi/edot/product-service/domain/products"
	"github.com/christiandwi/edot/product-service/entity"
)

type service struct {
	productRepo products.ProductsRepository
}

func NewProductService(productRepo products.ProductsRepository) Service {
	return &service{
		productRepo: productRepo,
	}
}

func (serv *service) GetProducts() (productList []entity.Products, err error) {
	productList, err = serv.productRepo.GetProducts()
	if err != nil {
		log.Print("error at getting products")
		return
	}

	return
}
