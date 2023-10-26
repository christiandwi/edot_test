package container

import (
	"github.com/christiandwi/edot/order-service/config"
	"github.com/christiandwi/edot/order-service/database"
	"github.com/christiandwi/edot/order-service/infrastructure/repo"
	"github.com/christiandwi/edot/order-service/usecase/order"
)

type Container struct {
	Config       *config.Config
	OrderService order.Service
}

func SetupContainer() (out Container) {
	cfg := config.SetupConfig()

	db := database.SetupDatabase(cfg)

	orderRepo := repo.SetupOrderRepo(db)

	//setup asynq package
	// asynqServer := asynq.NewServer(
	// 	asynq.RedisClientOpt{
	// 		Addr: cfg.Redis.Address,
	// 	},
	// 	asynq.Config{
	// 		Concurrency: cfg.Asynq.Concurrency,
	// 	},
	// )

	// mux := asynq.NewServeMux()
	//handler

	// go func() {
	// 	if err := asynqServer.Run(mux); err != nil {
	// 		log.Fatalf("could not run server: %v", err)
	// 	}
	// }()

	// client := asynq.NewClient(asynq.RedisClientOpt{
	// 	Addr: cfg.Redis.Address,
	// })
	// defer client.Close()

	//setup product
	orderService := order.NewOrderService(orderRepo)

	return Container{
		Config:       cfg,
		OrderService: orderService,
	}
}
