package server

import (
	"github.com/christiandwi/edot/product-service/response"
	"github.com/christiandwi/edot/product-service/usecase/product"
	"github.com/gin-gonic/gin"
)

type productHandler struct {
	service product.Service
}

func newProductHandler(service product.Service) *productHandler {
	return &productHandler{service: service}
}

func (g *productHandler) GetProducts() func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		products, err := g.service.GetProducts()
		if err != nil {
			response.SetResponse(ctx, false, 500, nil, err)
		}

		var productList []map[string]interface{}
		for _, v := range products {
			product := make(map[string]interface{})
			product["secure_id"] = v.SecureId
			product["product_name"] = v.ProductName
			product["stock_availability"] = v.StockAvailability
			productList = append(productList, product)
		}

		response.SetResponse(ctx, true, 200, productList, nil)
	}
}
