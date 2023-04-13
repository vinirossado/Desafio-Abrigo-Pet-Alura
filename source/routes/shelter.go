package routes

import (
	"abrigos/source/controllers"
	"abrigos/source/domain/enumerations"
	"abrigos/source/middleware"

	"github.com/gin-gonic/gin"
)

func bindShelterRouters(router *gin.Engine) {
	shelter := router.Group("/shelter")

	shelter.Use(middleware.JwtMiddleware().MiddlewareFunc())

	shelter.GET("", middleware.AuthorizationMiddleware(enumerations.ADMIN), controllers.FindShelters)
	shelter.GET("/paginated", middleware.AuthorizationMiddleware(enumerations.ADMIN), controllers.FindShelterPaginated)
	shelter.GET("/:id", middleware.AuthorizationMiddleware(enumerations.ADMIN), controllers.FindShelterById)
	shelter.POST("", middleware.AuthorizationMiddleware(enumerations.ADMIN), controllers.CreateShelter)
	shelter.PUT("/:id", middleware.AuthorizationMiddleware(enumerations.ADMIN), controllers.UpdateShelter)
	shelter.DELETE("/:id", middleware.AuthorizationMiddleware(enumerations.ADMIN), controllers.DeleteShelter)
	// shelter.PATCH("", AuthorizationMiddleware(enumerations.ADMIN),  controllers.)

}
