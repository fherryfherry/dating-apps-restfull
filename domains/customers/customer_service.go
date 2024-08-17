package customers

import (
	passCommon "booking-online/commons/password"
	"gorm.io/gorm"
	"time"
)

func InitCustomerService(DbCon *gorm.DB) CustomerService {
	return CustomerService{DbCon: DbCon}
}

type CustomerService struct {
	DbCon *gorm.DB
}

func (c *CustomerService) FindNonSwipeToday(customerId uint) *CustomerModel {
	result := &CustomerModel{}
	subQuery := c.DbCon.Table("swipes").Select("swipe_customers_id").Where("swipes.customers_id = ? AND DATE(swipe_at) = CURDATE()", customerId)
	query := c.DbCon.Model(&CustomerModel{}).Where("id NOT IN (?) and deleted_at IS NULL and id != ?", subQuery, customerId).Order("RAND()").First(&result)
	if query.Error != nil && query.Error.Error() != "record not found" {
		return nil
	}
	return result
}

func (c *CustomerService) FindByID(id uint) CustomerModel {
	customerData := CustomerModel{}
	result := c.DbCon.First(&customerData, "id = ? and deleted_at is null", id)
	if result.Error != nil && result.Error.Error() != "record not found" {
		return CustomerModel{}
	}
	return customerData
}

func (c *CustomerService) FindByUuid(uuid string) *CustomerModel {
	customerData := CustomerModel{}
	result := c.DbCon.First(&customerData, "customer_uuid = ? and deleted_at is null", uuid)
	if result.Error != nil && result.Error.Error() != "record not found" {
		return nil
	}
	return &customerData
}

func (c *CustomerService) FindByEmail(email string) CustomerModel {
	customerData := CustomerModel{}
	result := c.DbCon.First(&customerData, "email = ? and deleted_at is null", email)
	if result.Error != nil && result.Error.Error() != "record not found" {
		return CustomerModel{}
	}
	return customerData
}

func (c *CustomerService) CheckExistByEmail(email string) bool {
	customerData := CustomerModel{}
	result := c.DbCon.First(&customerData, "email = ? and deleted_at is null", email)
	if result.Error != nil && result.Error.Error() != "record not found" {
		return false
	}
	if customerData.Email != "" {
		return true
	} else {
		return false
	}
}

func (c *CustomerService) CreateCustomer(payload CreateCustomerPayload) (*CustomerModel, error) {
	newData := new(CustomerModel)
	newData.CustomerUUID = payload.UUID
	newData.FirstName = payload.FirstName
	newData.LastName = payload.LastName
	newData.ProfilePicture = payload.ProfilePicture
	newData.Bio = payload.Bio
	newData.Level = payload.Level
	newData.SwipeQuota = payload.SwipeQuota
	newData.Email = payload.Email
	passHashed, err := passCommon.HashPassword(payload.Password)
	if err != nil {
		return nil, err
	}
	newData.Password = passHashed

	result := c.DbCon.Create(&newData)
	if result.Error != nil {
		return nil, result.Error
	}

	return newData, nil
}

func (c *CustomerService) UpdateProfilePicture(id uint, profilePic string) error {
	updateData := c.FindByID(id)
	updateData.ProfilePicture = profilePic
	update := c.DbCon.Save(updateData)
	if update.Error != nil {
		return update.Error
	}
	return nil
}

func (c *CustomerService) UpdatePackageLevel(id uint, packageLevel string, packageQuota int64) error {
	updateData := c.FindByID(id)
	updateData.Level = packageLevel
	updateData.SwipeQuota = packageQuota
	updateData.PackageExpiry = time.Now().AddDate(0, 0, 30)
	if update := c.DbCon.Save(updateData); update.Error != nil {
		return update.Error
	}
	return nil
}
