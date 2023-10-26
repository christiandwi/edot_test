package container

import (
	"github.com/christiandwi/go-asynq/config"
	"github.com/christiandwi/go-asynq/database"
	"github.com/christiandwi/go-asynq/usecase/guest"
	"github.com/hibiken/asynq"
)

type Container struct {
	Config       *config.Config
	GuestService guest.Service
}

func SetupContainer() (out Container) {
	cfg := config.SetupConfig()

	db := database.SetupDatabase(cfg)

	//setup asynq package
	asynq := asynq.NewServer(
		asynq.RedisClientOpt{
			Addr: cfg.Redis.Address,
		},
		asynq.Config{
			Concurrency: cfg.Asynq.Concurrency,
		},
	)

	//setup guest
	guestService := guest.NewGuestService(db, asynq)

	return Container{
		Config:       cfg,
		GuestService: guestService,
	}
}
