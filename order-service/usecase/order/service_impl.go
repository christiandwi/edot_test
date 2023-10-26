package order

import (
	"log"

	"github.com/christiandwi/edot/order-service/domain/orders"
)

type service struct {
	orderRepo orders.OrderRepository
}

func NewOrderService(orderRepo orders.OrderRepository) Service {
	return &service{
		orderRepo: orderRepo,
	}
}

func (serv *service) Checkout(productId string, warehouseId string, userId string, amount float64) (err error) {
	err = serv.orderRepo.AddCheckoutLock(productId, userId, warehouseId, amount)
	if err != nil {
		log.Print("error at add checkout lock")
		return
	}

	return
}
