package main

import (
	"golang-gin-crud/internal/controllers"
	"golang-gin-crud/internal/initializer"

	"github.com/gin-gonic/gin"
)

func init() {
	initializer.LoadEnvValues()
	initializer.ConnectToDB()
}

func main() {

	r := gin.Default()
	r.POST("/employee/create", controllers.CreateEmployee)
	r.GET("/employees", controllers.GetAllEmployees)
	r.GET("/employee/:id", controllers.GetEmployee)
	r.PUT("/employee/:id", controllers.UpdateEmployee)
	r.DELETE("/employee/:id", controllers.DeleteEmployee)
	r.Run()
}
