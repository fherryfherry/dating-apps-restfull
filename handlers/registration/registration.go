package registration

import (
	"booking-online/domains/customers"
	"booking-online/domains/packages"
	"gorm.io/gorm"
)

type Registration struct {
	customerSrv customers.CustomerService
	packages    packages.PackageService
}

func InitHandler(dbConnection *gorm.DB) Registration {
	return Registration{
		customerSrv: customers.InitCustomerService(dbConnection),
		packages:    packages.InitPackageService(dbConnection),
	}
}
