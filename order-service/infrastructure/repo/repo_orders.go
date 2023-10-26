package repo

import (
	"fmt"
	"log"
	"time"

	"github.com/christiandwi/edot/order-service/database"
	"github.com/christiandwi/edot/order-service/domain/orders"
	"github.com/christiandwi/edot/order-service/entity"
	"gorm.io/gorm"
)

type orderRepo struct {
	db *database.Database
}

func SetupOrderRepo(db *database.Database) orders.OrderRepository {
	return &orderRepo{db: db}
}

func (o orderRepo) AddCheckoutLock(productId string, userId string, warehouseId string, amount float64) (err error) {
	//get product id from secureid
	var product entity.Products
	if err = o.db.Where("secure_id = ?", productId).First(&product).Error; err != nil {
		log.Print("error on getting product")
		return
	}

	//get user id from secureid
	var user entity.Users
	if err = o.db.Where("secure_id = ?", userId).First(&user).Error; err != nil {
		log.Print("error on getting user")
		return
	}

	//get warehouse id from secureid
	var warehouse entity.Warehouses
	if err = o.db.Where("secure_id = ?", warehouseId).First(&warehouse).Error; err != nil {
		log.Print("error on getting warehouse")
		return
	}

	if amount > product.StockAvailability {
		err = fmt.Errorf("stock is unavailable")
		return err
	}

	//this process will be better if use a queue system

	tx := o.db.Begin()

	if err = tx.Model(&product).Update("stock_availability", gorm.Expr("stock_availability - ?", amount)).Error; err != nil {
		log.Print("error on deduct stock")
		return
	}

	var order entity.Orders
	order.UserId = user.ID
	order.WarehouseId = warehouse.ID
	order.ProductId = product.ID
	order.Amount = amount
	order.ExpiredAt = time.Now().Add(1 * time.Second)

	if err = tx.Create(&order).Error; err != nil {
		return
	}

	tx.Commit()

	//this is a workaround, not a best practice as it will spawn too many goroutines if the system grows, the best practice would be to do a background job every minute
	go func() {
		time.Sleep(2 * time.Second)

		if err = o.db.Model(&product).Update("stock_availability", gorm.Expr("stock_availability + ?", order.Amount)).Error; err != nil {
			log.Print("error on returning stock")
			return
		}

		o.db.Delete(order)
	}()

	return
}
