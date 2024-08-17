package profile

import (
	"booking-online/domains/customers"
	"gorm.io/gorm"
)

type Profile struct {
	customer customers.CustomerService
}

func InitHandler(dbConnection *gorm.DB) Profile {
	return Profile{
		customer: customers.InitCustomerService(dbConnection),
	}
}
