package products

import "github.com/christiandwi/edot/product-service/entity"

type ProductsRepository interface {
	GetProducts() (products []entity.Products, err error)
}
