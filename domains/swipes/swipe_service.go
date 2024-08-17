package swipes

import (
	"gorm.io/gorm"
	"time"
)

func InitSwipeService(DbCon *gorm.DB) SwipeService {
	return SwipeService{DbCon: DbCon}
}

type SwipeService struct {
	DbCon *gorm.DB
}

func (c *SwipeService) GetSwipeToday(customerID uint) int64 {
	var count int64
	query := c.DbCon.Model(&SwipeModel{}).Where("customers_id = ? and DATE(swipe_at) = CURDATE()", customerID).Count(&count)
	if query.Error != nil && query.Error.Error() != "record not found" {
		return 0
	}
	return count
}

func (c *SwipeService) CheckSwipe(customerID uint, swipeCustomerID uint) bool {
	var count int64
	query := c.DbCon.Model(&SwipeModel{}).Where("customers_id = ? and swipe_customers_id = ? and DATE(swipe_at) = CURDATE()", customerID, swipeCustomerID).Count(&count)
	if query.Error != nil && query.Error.Error() != "record not found" {
		return false
	}
	return count > 0
}

func (c *SwipeService) AddSwipe(customerId uint, swipeCustomerId uint, swipeType string) error {
	add := SwipeModel{
		CustomerID:      customerId,
		SwipeCustomerID: swipeCustomerId,
		SwipeAt:         time.Now(),
		SwipeType:       swipeType,
	}
	query := c.DbCon.Create(&add)
	if query.Error != nil {
		return query.Error
	}
	return nil
}
