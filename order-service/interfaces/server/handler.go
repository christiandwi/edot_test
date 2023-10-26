package server

import (
	"github.com/christiandwi/edot/order-service/interfaces/container"
)

type handler struct {
	orderHandler *orderHandler
}

func setupHandler(container container.Container) *handler {
	orderHandler := newOrderHandler(container.OrderService)
	return &handler{
		orderHandler: orderHandler,
	}
}
