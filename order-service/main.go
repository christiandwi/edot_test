package main

import (
	"github.com/christiandwi/edot/order-service/interfaces/container"
	"github.com/christiandwi/edot/order-service/interfaces/server"
)

func main() {
	server.Start(container.SetupContainer())
}
