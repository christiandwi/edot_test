package order

type Service interface {
	Checkout(productId string, warehouseId string, userId string, amount float64) (err error)
}
