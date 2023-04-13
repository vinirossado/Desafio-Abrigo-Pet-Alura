package service

import (
	"abrigos/source/domain/entities"
	"abrigos/source/domain/request"
	"abrigos/source/repository"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func FindUsers() {

}

func FindUserById(c *gin.Context) {

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
