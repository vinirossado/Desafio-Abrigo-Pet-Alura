package entities

import "gorm.io/gorm"

type Shelter struct {
	gorm.Model
	Name         string `gorm:"column:name;unique"`
	City         string `gorm:"column:city"`
	ZipCode      string `gorm:"column:zip_code"`
	Address      string `gorm:"column:address"`
	Number       string `gorm:"column:number"`
	Neighborhood string `gorm:"column:neighborhood"`
	Complement   string `gorm:"column:complement"`

	ResponsibleID uint `gorm:"column:responsible_id"`
	User          User `gorm:"foreignKey:ResponsibleID"`

	Pets []Pet `gorm:"foreignkey:ShelterID"`
}
