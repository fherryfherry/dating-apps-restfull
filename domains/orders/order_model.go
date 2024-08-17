package orders

import (
	"gorm.io/gorm"
)

type OrderModel struct {
	gorm.Model
	OrderNo       string  `gorm:"column:order_no;type:varchar(255)"`
	CustomerID    uint    `gorm:"column:customer_id"`
	CustomerName  string  `gorm:"column:customer_name;type:varchar(255)"`
	CustomerEmail string  `gorm:"column:customer_email;type:varchar(255)"`
	PackageID     uint    `gorm:"column:packages_id"`
	PackageTitle  string  `gorm:"column:packages_title;type:varchar(255)"`
	PackageQuota  int64   `gorm:"column:packages_quota"`
	GrandTotal    float32 `gorm:"column:grand_total"`
	OrderStatus   string  `gorm:"column:order_status;type:varchar(55)"`
	PaymentStatus string  `gorm:"column:payment_status;type:varchar(55)"`
}

type Tabler interface {
	TableName() string
}

func (OrderModel) TableName() string {
	return "orders"
}
