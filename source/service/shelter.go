package service

import (
	"abrigos/source/domain/entities"
	"abrigos/source/domain/exception"
	"abrigos/source/domain/request"
	"abrigos/source/domain/responses"
	"abrigos/source/middleware"
	"abrigos/source/repository"

	"errors"
	"fmt"

	"gorm.io/gorm"
)

func FindShelters() []responses.ShelterResponse {
	shelters, err := repository.FindShelters()

	if err != nil {
		exception.ThrowInternalServerException(
			fmt.Sprintf("Error while trying to get all shelters with error: %s", err),
		)
	}

	sheltersResponse := []responses.ShelterResponse{}

	for _, shelter := range shelters {
		sheltersResponse = append(sheltersResponse, *MapToShelterResponse(&shelter))
	}

	return sheltersResponse
}

func FindShelterById(id int) *responses.ShelterResponse {
	shelter, err := repository.FindShelterById(id)

	if err != nil {
		exception.ThrowNotFoundException(fmt.Sprintf("shelter with %d was not found", id))
	}

	return MapToShelterResponse(shelter)
}

func CreateShelter(request *request.ShelterRequest) {

	repository.UsingTransactional(func(tx *repository.TransactionalOperation) error {
		exists := repository.ExistsShelterByName(request.Name)

		if exists {
			return exception.BadRequestException(
				fmt.Sprintf("Name %s already exists", request.Name),
			)
		}

		claimID := middleware.GlobalClaims["id"].(float64)

		shelter := entities.Shelter{
			Name:          request.Name,
			City:          request.City,
			ZipCode:       request.ZipCode,
			Address:       request.Address,
			Neighborhood:  request.Neighborhood,
			Number:        request.Number,
			ResponsibleID: uint(claimID),
			Complement:    request.Complement,
		}

		if err := repository.CreateShelter(&shelter, tx); err != nil {
			return exception.InternalServerException(
				fmt.Sprintf("Error while trying to insert new Shelter with error: %s", err),
			)
		}

		return nil
	})
}

func UpdateShelter(request *request.ShelterRequest, id int) {
	repository.UsingTransactional(func(tx *repository.TransactionalOperation) error {
		shelter, err := repository.FindShelterById(id)

		if err != nil {
			return exception.NotFoundException(
				fmt.Sprintf("Shelter with id {%d} was not found", id),
			)
		}

		shelter.Name = request.Name
		shelter.City = request.City
		shelter.ZipCode = request.ZipCode
		shelter.Address = request.Address
		shelter.Neighborhood = request.Neighborhood
		shelter.Number = request.Number
		shelter.ResponsibleID = 1
		shelter.Complement = request.Complement

		if err := repository.UpdateShelter(shelter, tx); err != nil {
			return exception.InternalServerException(
				fmt.Sprintf("Error ocurred while trying to update shelter with error: %s", err))
		}
		return nil
	})
}

func DeleteShelter(id int) {
	repository.UsingTransactional(func(tx *repository.TransactionalOperation) error {
		if err := repository.DeleteShelter(id, tx); err != nil {

			if errors.Is(err, gorm.ErrRecordNotFound) {
				return exception.NotFoundException(
					fmt.Sprintf("Shelter with id {%d} not found", id))
			}

			return exception.InternalServerException(
				fmt.Sprintf("Error ocurred while trying to delete new shelter with error: %s", err))
		}
		return nil
	})
}

func MapToShelterResponse(user *entities.Shelter) (response *responses.ShelterResponse) {

	return &responses.ShelterResponse{
		Name:          user.Name,
		City:          user.City,
		ZipCode:       user.ZipCode,
		Address:       user.Address,
		Neighborhood:  user.Neighborhood,
		Number:        user.Number,
		ResponsibleID: user.ResponsibleID,
		Complement:    user.Complement,
	}
}
