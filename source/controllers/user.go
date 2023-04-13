package controllers

import (
	"abrigos/source/domain/request"
	"abrigos/source/service"
	"abrigos/source/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func FindUsers(c *gin.Context) {
	users := service.FindUsers()
	c.JSON(http.StatusOK, users)
}

func FindUserById(c *gin.Context) {

}

func CreateUser(c *gin.Context) {
	user := request.UserRequest{}
	utils.ReadRequestBody(c, &user)
	service.CreateUser(&user)
	c.Status(http.StatusOK)

}

func UpdateUser(c *gin.Context) {

}

func DeleteUser(c *gin.Context) {

}
