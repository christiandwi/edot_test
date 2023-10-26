package main

import (
	"github.com/christiandwi/go-asynq/interfaces/container"
	"github.com/christiandwi/go-asynq/interfaces/server"
)

func main() {
	server.Start(container.SetupContainer())
}
