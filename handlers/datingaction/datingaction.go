package datingaction

import (
	"booking-online/domains/customers"
	"booking-online/domains/swipes"
	"gorm.io/gorm"
)

type DatingAction struct {
	swipe    swipes.SwipeService
	customer customers.CustomerService
}

func InitHandler(dbConnection *gorm.DB) DatingAction {
	return DatingAction{
		swipe:    swipes.InitSwipeService(dbConnection),
		customer: customers.InitCustomerService(dbConnection),
	}
}
