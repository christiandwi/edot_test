package product

import "github.com/christiandwi/edot/product-service/entity"

type Service interface {
	GetProducts() (productList []entity.Products, err error)
}
