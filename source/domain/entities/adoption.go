package entities

import "gorm.io/gorm"

type Adoption struct {
	gorm.Model
	PetID  uint `gorm:"column:pet_id"`
	UserID uint `gorm:"column:user_id"`

	Pet  Pet  `gorm:"foreignkey:PetID"`
	User User `gorm:"foreignkey:UserID"`
}
