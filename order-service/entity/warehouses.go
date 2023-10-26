package entity

import (
	"github.com/christiandwi/edot/order-service/constant"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Warehouses struct {
	ID       int64  `gorm:"column:id;primary_key"`
	SecureId string `gorm:"column:secure_id"`
	ShopId   int64
}

func (Warehouses) TableName() string {
	return constant.EntityWarehouses
}

func (w *Warehouses) BeforeCreate(scope *gorm.DB) (err error) {
	w.SecureId = uuid.NewString()

	return nil
}
