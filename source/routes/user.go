package routes

import (
	"abrigos/source/controllers"
	"abrigos/source/domain/enumerations"
	"abrigos/source/middleware"

	"github.com/gin-gonic/gin"
)

func bindUserRoutes(router *gin.Engine) {

	users := router.Group("/users")
	// users.Use(middleware.JwtMiddleware().MiddlewareFunc())

	users.GET("", controllers.FindUsers)
	users.GET("/:id", controllers.FindUserById)
	users.POST("", controllers.CreateUser)
	users.PUT("/:id", middleware.AuthorizationMiddleware(enumerations.SUPERVISOR), controllers.UpdateUser)
	users.PATCH("/:id", middleware.AuthorizationMiddleware(enumerations.ADMIN), controllers.DeleteUser)

}
