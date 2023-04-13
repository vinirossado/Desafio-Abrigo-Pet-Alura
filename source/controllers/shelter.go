package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func FindShelters(c *gin.Context) {
	// shelters := services.FindShelters()
	c.JSON(http.StatusOK, nil)
}

func FindShelterPaginated(c *gin.Context) {

}

func FindShelterById(c *gin.Context) {

}

func CreateShelter(c *gin.Context) {

}

func UpdateShelter(c *gin.Context) {

}

func DeleteShelter(c *gin.Context) {

}
