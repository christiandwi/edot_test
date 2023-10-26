package database

import (
	"fmt"
	"log"

	"github.com/christiandwi/edot/order-service/config"
	"github.com/christiandwi/edot/order-service/constant"
	"github.com/christiandwi/edot/order-service/entity"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type Database struct {
	*gorm.DB
}

func SetupDatabase(config *config.Config) *Database {
	var (
		db  *gorm.DB
		err error
	)

	url := fmt.Sprintf(config.DB.Datasource)
	dialect := fmt.Sprintf(config.DB.Dialect)

	if dialect == constant.MysqlDialect {
		db, err = gorm.Open(mysql.Open(url), &gorm.Config{
			Logger: logger.Default.LogMode(logger.Info),
		})
	}
	if err != nil {
		log.Print("error at init db")
		panic(err)
	}

	sqldb, err := db.DB()
	if err != nil {
		log.Print("error at init sqldb instance")
		panic(err)
	}

	sqldb.SetMaxIdleConns(config.DB.MaxIdleConns)
	sqldb.SetMaxOpenConns(config.DB.MaxOpenConns)

	db.AutoMigrate(&entity.Orders{}, &entity.Warehouses{})

	return &Database{db}
}
