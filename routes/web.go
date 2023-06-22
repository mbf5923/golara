package routes

import (
	"github.com/gin-gonic/gin"
	"gorm-test/app/http/controllers"
	"gorm-test/app/http/middleware"
)

func InitialRoutes(route *gin.Engine) {

	groupRoute := route.Group("api/v1")
	groupRoute.GET("/", func(c *gin.Context) {
		c.JSON(200, "GOLARA API V1")
	})
	userRepo := controllers.New()
	groupRoute.POST("/users", userRepo.CreateUser)
	groupRoute.POST("/users/login", userRepo.Login)
	groupRoute.GET("/users", userRepo.GetUsers)
	groupRoute.GET("/users/:id", userRepo.GetUser)
	groupRoute.PUT("/users/:id", userRepo.UpdateUser)
	groupRoute.DELETE("/users/:id", userRepo.DeleteUser)

	bookRepo := controllers.NewBookRepo()
	groupRoute.POST("/books", middleware.Auth(), bookRepo.CreateBook)
	groupRoute.GET("/books", middleware.Auth(), bookRepo.Books)
	groupRoute.PUT("/books/:id", middleware.Auth(), bookRepo.UpdateBook)
}
