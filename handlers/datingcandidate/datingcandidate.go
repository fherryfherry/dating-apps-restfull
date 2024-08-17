package datingcandidate

import (
	"booking-online/domains/customers"
	"booking-online/domains/swipes"
	"gorm.io/gorm"
)

type DatingCandidate struct {
	customer customers.CustomerService
	swipe    swipes.SwipeService
}

func InitHandler(dbConnection *gorm.DB) DatingCandidate {
	return DatingCandidate{
		customer: customers.InitCustomerService(dbConnection),
		swipe:    swipes.InitSwipeService(dbConnection),
	}
}
