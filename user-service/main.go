package main

import (
	"github.com/christiandwi/edot/user-service/interfaces/container"
	"github.com/christiandwi/edot/user-service/interfaces/server"
)

func main() {
	server.Start(container.SetupContainer())
}
