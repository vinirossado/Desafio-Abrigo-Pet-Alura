package entities

import "gorm.io/gorm"

type Pet struct {
	gorm.Model
	Name            string `gorm:"column:name"`
	Description     string `gorm:"column:description"`
	Adopted         bool   `gorm:"column:adopted"`
	Age             string `gorm:"column:age"`
	Size            string `gorm:"column:size"`
	Image           string `gorm:"column:image"`
	Characteristics string `gorm:"column:characteristics"`
	ShelterID       uint   `gorm:"column:shelter_id"`

	Shelter Shelter `gorm:"foreignKey:ShelterID"`
}
