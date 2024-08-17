package swipes

import (
	"gorm.io/gorm"
	"time"
)

type SwipeModel struct {
	gorm.Model
	CustomerID      uint      `gorm:"column:customers_id"`
	SwipeCustomerID uint      `gorm:"column:swipe_customers_id"`
	SwipeAt         time.Time `gorm:"column:swipe_at"`
	SwipeType       string    `gorm:"column:swipe_type;type:varchar(55)"`
}

type Tabler interface {
	TableName() string
}

func (SwipeModel) TableName() string {
	return "swipes"
}
