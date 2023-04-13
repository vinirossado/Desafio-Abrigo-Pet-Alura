package routes

import (
	"abrigos/source/middleware"

	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	router := gin.New()
	// bindSwagger(router)
	bindMiddlewares(router)
	bindShelterRouters(router)
	bindActuatorsRoutes(router)
	bindUserRoutes(router)

	return router
}

func bindMiddlewares(router *gin.Engine) {
	router.Use(gin.Logger())
	router.Use(gin.CustomRecovery(middleware.ExceptionMiddleware))
}
