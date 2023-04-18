package controllers

import (
	"abrigos/source/domain/request"
	"abrigos/source/service"
	"abrigos/source/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func FindShelters(c *gin.Context) {
	shelters := service.FindShelters()
	c.JSON(http.StatusOK, shelters)
}

func FindShelterPaginated(c *gin.Context) {

}

func FindShelterById(c *gin.Context) {
	id := utils.ConvertToInt(c.Params.ByName("id"))
	user := service.FindShelterById(id)
	c.JSON(http.StatusOK, user)
}

func CreateShelter(c *gin.Context) {
	shelter := request.ShelterRequest{}
	utils.ReadRequestBody(c, &shelter)
	service.CreateShelter(&shelter)
	c.Status(http.StatusOK)
}

func UpdateShelter(c *gin.Context) {
	updateShelterRequest := request.ShelterRequest{}
	utils.ReadRequestBody(c, &updateShelterRequest)

	id := utils.ConvertToInt(c.Params.ByName("id"))
	service.UpdateShelter(&updateShelterRequest, id)
	c.Status(http.StatusOK)
}

func DeleteShelter(c *gin.Context) {
	id := utils.ConvertToInt(c.Params.ByName("id"))
	service.DeleteShelter(id)
	c.Status(http.StatusOK)
}
