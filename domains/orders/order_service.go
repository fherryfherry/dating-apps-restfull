package orders

import (
	"fmt"
	"gorm.io/gorm"
)

func InitOrderService(DbCon *gorm.DB) OrderService {
	return OrderService{DbCon: DbCon}
}

type OrderService struct {
	DbCon *gorm.DB
}

func (c *OrderService) GetListByCustomer(customerId uint) []OrderModel {
	var orderList []OrderModel
	result := c.DbCon.Order("id desc").Find(&orderList, "customer_id = ? and deleted_at is null", customerId)
	if result.Error != nil && result.Error.Error() != "record not found" {
		return nil
	}
	return orderList
}

func (c *OrderService) CountOrder() int64 {
	var total int64
	c.DbCon.Table("orders").Count(&total)
	return total
}

func (c *OrderService) CreateOrder(payload CreateOrderPayload) (string, error) {
	total := c.CountOrder() + 1
	orderNo := "ORD" + fmt.Sprintf("%03d", total)
	orderMaster := new(OrderModel)
	orderMaster.OrderNo = orderNo
	orderMaster.OrderStatus = "COMPLETED"
	orderMaster.PaymentStatus = "PAID"
	orderMaster.CustomerID = payload.CustomerID
	orderMaster.CustomerName = payload.CustomerName
	orderMaster.CustomerEmail = payload.CustomerEmail
	orderMaster.PackageID = payload.PackageID
	orderMaster.PackageTitle = payload.PackageTitle
	orderMaster.PackageQuota = payload.PackageQuota
	orderMaster.GrandTotal = payload.PackagePrice
	orderCreateResult := c.DbCon.Create(&orderMaster)
	if orderCreateResult.Error != nil {
		return "", orderCreateResult.Error
	}

	return orderNo, nil
}
