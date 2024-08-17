package packages

import (
	"gorm.io/gorm"
)

func InitPackageService(DbCon *gorm.DB) PackageService {
	return PackageService{DbCon: DbCon}
}

type PackageService struct {
	DbCon *gorm.DB
}

func (c *PackageService) FindByCode(code string) *PackageModel {
	packModel := &PackageModel{}
	result := c.DbCon.First(&packModel, "code = ?", code)
	if result.Error != nil && result.Error.Error() != "record not found" {
		return nil
	}
	return packModel
}

func (c *PackageService) GetList() []PackageModel {
	packModels := []PackageModel{}
	result := c.DbCon.Find(&packModels)
	if result.Error != nil && result.Error.Error() != "record not found" {
		return nil
	}
	return packModels
}
