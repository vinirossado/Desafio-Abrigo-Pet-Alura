package utils

import "github.com/gin-gonic/gin"

func ReadRequestBody(c *gin.Context, requestBody interface{}) {

	err := c.ShouldBindJSON(&requestBody)

	if err != nil {
		//Tratar erro
	}

}
