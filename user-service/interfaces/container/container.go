package container

import (
	"github.com/christiandwi/edot/user-service/config"
	"github.com/christiandwi/edot/user-service/database"
	"github.com/christiandwi/edot/user-service/infrastructure/repo"
	"github.com/christiandwi/edot/user-service/usecase/guest"
)

type Container struct {
	Config       *config.Config
	GuestService guest.Service
}

func SetupContainer() (out Container) {
	cfg := config.SetupConfig()

	db := database.SetupDatabase(cfg)

	userRepo := repo.SetupUserRepo(db)

	//setup asynq package
	// asynq := asynq.NewServer(
	// 	asynq.RedisClientOpt{
	// 		Addr: cfg.Redis.Address,
	// 	},
	// 	asynq.Config{
	// 		Concurrency: cfg.Asynq.Concurrency,
	// 	},
	// )

	//setup guest
	guestService := guest.NewGuestService(userRepo)

	return Container{
		Config:       cfg,
		GuestService: guestService,
	}
}
