package customers

import (
	"gorm.io/gorm"
	"time"
)

type CustomerModel struct {
	gorm.Model
	CustomerUUID   string    `gorm:"column:customer_uuid;type:varchar(255)"`
	FirstName      string    `gorm:"column:first_name;type:varchar(255)"`
	LastName       string    `gorm:"column:last_name;type:varchar(255)"`
	Bio            string    `gorm:"column:bio;type:varchar(500)"`
	Email          string    `gorm:"column:email;type:varchar(255)"`
	Password       string    `gorm:"column:password;type:varchar(255)"`
	Level          string    `gorm:"column:level;type:varchar(255)"`
	SwipeQuota     int64     `gorm:"column:swipe_quota"`
	PackageExpiry  time.Time `gorm:"column:package_expiry"`
	ProfilePicture string    `gorm:"column:profile_picture;type:varchar(255)"`
}

type Tabler interface {
	TableName() string
}

// TableName overrides the table name used by User to `profiles`
func (CustomerModel) TableName() string {
	return "customers"
}
