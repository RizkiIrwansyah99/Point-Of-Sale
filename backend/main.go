package main

import (
	"github.com/arioprima/blog_web/config"
	"github.com/arioprima/blog_web/controller"
	"github.com/arioprima/blog_web/repository"
	"github.com/arioprima/blog_web/service"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	r := gin.Default()

	// connect to database
	db := config.ConnectionDB()

	// validate
	validate := validator.New()

	//repository
	userRepository := repository.NewUserRepositoryImpl(db)

	//service
	userService := service.NewUserServiceImpl(userRepository, db, validate)

	//controller
	userController := controller.NewUserController(userService)

	// initialize
	r.POST("/api/users", userController.Create)
	r.PUT("/api/users", userController.Update)
	r.DELETE("/api/users/:id", userController.Delete)
	r.GET("/api/users/:id", userController.FindById)
	r.GET("/api/users/username/:username", userController.FindByUserName)
	r.GET("/api/users/email/:email", userController.FindByEmail)
	r.GET("/api/users", userController.FindAll)

	err := r.Run(":8080")
	if err != nil {
		panic(err)
	}
}
