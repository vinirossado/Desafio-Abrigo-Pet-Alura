package service

import (
	"abrigos/source/domain/entities"
	"abrigos/source/domain/request"
	"abrigos/source/repository"

	"github.com/gin-gonic/gin"
)

func FindUsers() {

}

func FindUserById(c *gin.Context) {

}

func CreateUser(request *request.UserRequest) {

	user := entities.User{
		Name:     request.Name,
		Username: request.Username,
		Password: request.Password,
		Role:     request.Role,
	}
	repository.CreateUser(&user)
}

func UpdateUser(c *gin.Context) {

}

func DeleteUser(c *gin.Context) {

}
