package service

import (
	"abrigos/source/domain/entities"
	"abrigos/source/domain/exception"
	"abrigos/source/domain/request"
	"abrigos/source/domain/responses"
	"abrigos/source/repository"
	"errors"
	"fmt"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func FindUsers() []responses.UserResponse {
	users, err := repository.FindUsers()

	if err != nil {
		exception.ThrowInternalServerException(
			fmt.Sprintf("Error while trying to get all tutors with error: %s", err),
		)
	}

	usersResponse := []responses.UserResponse{}

	for _, user := range users {
		usersResponse = append(usersResponse, *MapToUserResponse(&user))
	}

	return usersResponse
}

func FindUserById(id int) *responses.UserResponse {
	user, err := repository.FindUserById(id)

	if err != nil {
		exception.ThrowNotFoundException(fmt.Sprintf("User with %d was not found", id))
	}

	return MapToUserResponse(user)
}

func CreateUser(request *request.UserRequest) {

	hash, err := bcrypt.GenerateFromPassword([]byte(request.Password), bcrypt.DefaultCost)

	if err != nil {
		// Tratar erro
	}

	user := entities.User{
		Name:     request.Name,
		Username: request.Username,
		Password: string(hash),
		Role:     request.Role,
	}
	repository.CreateUser(&user)
}

func UpdateUser(request *request.UserRequest, id int) {
	repository.UsingTransactional(func(tx *repository.TransactionalOperation) error {
		user, err := repository.FindUserById(id)

		if err != nil {
			return exception.NotFoundException(
				fmt.Sprintf("User with id {%d} not found", id))
		}

		user.Name = request.Name
		user.Username = request.Username
		user.Password = request.Password

		if err := repository.UpdateUser(user, tx); err != nil {
			return exception.InternalServerException(
				fmt.Sprintf("Error ocurred while trying to update new user with error: %s", err))
		}
		return nil
	})
}

func DeleteUser(id int) {
	repository.UsingTransactional(func(tx *repository.TransactionalOperation) error {
		if err := repository.DeleteUser(id, tx); err != nil {

			if errors.Is(err, gorm.ErrRecordNotFound) {
				return exception.NotFoundException(
					fmt.Sprintf("User with id {%d} not found", id))
			}

			return exception.InternalServerException(
				fmt.Sprintf("Error ocurred while trying to delete new user with error: %s", err))
		}
		return nil
	})
}

func MapToUserResponse(user *entities.User) (response *responses.UserResponse) {

	return &responses.UserResponse{
		ID:       user.ID,
		Name:     user.Name,
		Username: user.Username,
		Role:     user.Role,
	}
}
