package service

import (
	"abrigos/source/domain/entities"
	"abrigos/source/domain/exception"
	"abrigos/source/domain/request"
	"abrigos/source/domain/responses"
	"abrigos/source/repository"
	"fmt"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
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

func UpdateUser(c *gin.Context) {

}

func DeleteUser(c *gin.Context) {

}

func MapToUserResponse(user *entities.User) (response *responses.UserResponse) {

	return &responses.UserResponse{
		ID:       user.ID,
		Name:     user.Name,
		Username: user.Username,
		Role:     user.Role,
	}
}
