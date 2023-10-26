package guest

import (
	"github.com/christiandwi/go-asynq/database"
	"github.com/hibiken/asynq"
)

type service struct {
	db    *database.Database
	asynq *asynq.Server
}

func NewGuestService(db *database.Database, asynq *asynq.Server) Service {
	return &service{
		db:    db,
		asynq: asynq,
	}
}

func (serv *service) HelloService(name string) (text map[string]interface{}) {
	return map[string]interface{}{"text": "hello, " + name}
}
