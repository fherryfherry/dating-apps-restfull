package order

import (
	"booking-online/domains/customers"
	"booking-online/domains/orders"
	"booking-online/domains/packages"
	"gorm.io/gorm"
)

type Order struct {
	orderSrv orders.OrderService
	packages packages.PackageService
	customer customers.CustomerService
}

func InitHandler(dbConnection *gorm.DB) Order {
	return Order{
		orderSrv: orders.InitOrderService(dbConnection),
		packages: packages.InitPackageService(dbConnection),
		customer: customers.InitCustomerService(dbConnection),
	}
}
