package repository

import (
	"abrigos/source/domain/entities"

	"gorm.io/gorm"
)

func FindShelters(tx ...*TransactionalOperation) ([]entities.Shelter, error) {
	shelters := []entities.Shelter{}
	return shelters, WithTransaction(tx).Find(&shelters).Error
}

func FindShelterByName(name string, tx ...*TransactionalOperation) (*entities.Shelter, error) {
	shelter := &entities.Shelter{}
	return shelter, WithTransaction(tx).Where("name = ?", name).First(shelter).Error
}

func ExistsShelterByName(name string, tx ...*TransactionalOperation) bool {
	dbResult := WithTransaction(tx).Where("name = ?", name).Find(&entities.Shelter{})
	return dbResult.RowsAffected > 0
}

func FindShelterById(id int, tx ...*TransactionalOperation) (*entities.Shelter, error) {
	shelter := &entities.Shelter{}
	return shelter, WithTransaction(tx).Where("id = ?", id).First(shelter).Error
}

func CreateShelter(shelter *entities.Shelter, tx ...*TransactionalOperation) error {
	return WithTransaction(tx).Create(shelter).Error
}

func UpdateShelter(shelter *entities.Shelter, tx ...*TransactionalOperation) error {
	return WithTransaction(tx).Updates(&shelter).Error
}

func DeleteShelter(id int, tx ...*TransactionalOperation) error {
	dbResult := WithTransaction(tx).Delete(&entities.Shelter{Model: gorm.Model{ID: uint(id)}})

	if dbResult.RowsAffected < 1 {
		return gorm.ErrRecordNotFound
	}

	return dbResult.Error
}
