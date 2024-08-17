package quotapackage

import (
	"booking-online/domains/customers"
	"booking-online/domains/packages"
	"gorm.io/gorm"
)

type QuotaPackage struct {
	customer customers.CustomerService
	packages packages.PackageService
}

func InitHandler(dbConnection *gorm.DB) QuotaPackage {
	return QuotaPackage{
		customer: customers.InitCustomerService(dbConnection),
		packages: packages.InitPackageService(dbConnection),
	}
}
