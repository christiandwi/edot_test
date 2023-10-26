package entity

import (
	"time"

	"github.com/christiandwi/edot/order-service/constant"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Orders struct {
	ID          int64     `gorm:"column:id;primary_key"`
	SecureId    string    `gorm:"column:secure_id"`
	UserId      int64     `gorm:"column:user_id"`
	WarehouseId int64     `gorm:"column:warehouse_id"`
	ProductId   int64     `gorm:"column:product_id"`
	Amount      float64   `gorm:"column:amount"`
	ExpiredAt   time.Time `gorm:"column:expired_at"`
}

func (Orders) TableName() string {
	return constant.EntityOrders
}

func (o *Orders) BeforeCreate(scope *gorm.DB) (err error) {
	o.SecureId = uuid.NewString()

	return nil
}
