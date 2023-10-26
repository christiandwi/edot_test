package entity

import (
	"github.com/christiandwi/edot/product-service/constant"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Products struct {
	ID                int64  `gorm:"column:id;primary_key"`
	SecureId          string `gorm:"column:secure_id"`
	WarehouseId       int64
	ProductName       string
	StockAvailability float64
}

func (Products) TableName() string {
	return constant.EntityProducts
}

func (u *Products) BeforeCreate(scope *gorm.DB) (err error) {
	u.SecureId = uuid.NewString()

	return nil
}
