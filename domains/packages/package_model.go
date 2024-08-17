package packages

import "gorm.io/gorm"

type PackageModel struct {
	gorm.Model
	Code  string  `gorm:"column:code;type:varchar(255)"`
	Title string  `gorm:"column:title;type:varchar(255)"`
	Price float32 `gorm:"column:price"`
	Quota int64   `gorm:"column:quota"`
}

type Tabler interface {
	TableName() string
}

func (PackageModel) TableName() string {
	return "packages"
}
