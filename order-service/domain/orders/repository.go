package orders

type OrderRepository interface {
	AddCheckoutLock(productId string, userId string, warehouseId string, amount float64) (err error)
}
