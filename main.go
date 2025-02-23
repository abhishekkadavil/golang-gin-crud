package main

import (
	_ "golang-gin-crud/docs"

	"golang-gin-crud/internal/controllers"
	"golang-gin-crud/internal/initializer"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// gin-swagger middleware
// swagger embed files

func init() {
	initializer.LoadEnvValues()
	initializer.ConnectToDB()
}

// @title Golang Gin CRUD API
// @version	1.0
// @description Simple CRUD operation APIs in Golang Gin

// @contact.name   Abhishek Kadavil
// @contact.url    https://github.com/abhishekkadavil/golang-gin-crud
// @contact.email  abhishek.kadavil@gmail.com

// @host 	localhost:3000
// @BasePath /api
func main() {

	r := gin.Default()

	// add swagger
	r.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	r.POST("/employee/create", controllers.CreateEmployee)
	r.GET("/employees", controllers.GetAllEmployees)
	r.GET("/employee/:id", controllers.GetEmployee)
	r.PUT("/employee/:id", controllers.UpdateEmployee)
	r.DELETE("/employee/:id", controllers.DeleteEmployee)
	r.Run()
}
