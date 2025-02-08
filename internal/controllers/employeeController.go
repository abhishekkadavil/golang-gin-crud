package controllers

import (
	"golang-gin-crud/internal/initializer"
	"golang-gin-crud/internal/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateEmployee(c *gin.Context) {

	// Get data from req body
	var reqBody struct {
		Employee_id          uint8 `gorm:"primary_key"`
		Employee_name        string
		Employee_age         uint8
		Employee_salary      float32
		Employee_Designation string
	}

	c.Bind(&reqBody)

	// create employee
	employee := models.Employees{
		Employee_id:          reqBody.Employee_id,
		Employee_name:        reqBody.Employee_name,
		Employee_age:         reqBody.Employee_age,
		Employee_salary:      reqBody.Employee_salary,
		Employee_Designation: reqBody.Employee_Designation}
	result := initializer.DB.Create(&employee)

	if result.Error != nil {
		c.Status(http.StatusUnprocessableEntity)
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"Employee": employee,
	})
}

func GetAllEmployees(c *gin.Context) {

	// Get employee
	var employees []models.Employees
	result := initializer.DB.Find(&employees)

	if result.Error != nil {
		c.Status(http.StatusUnprocessableEntity)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"Employees": employees,
	})

}

func GetEmployee(c *gin.Context) {

	// Get id from url
	id := c.Param("id")

	// Get employee
	var employee models.Employees
	result := initializer.DB.First(&employee, id)

	if result.Error != nil {
		c.Status(http.StatusUnprocessableEntity)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"Employee": employee,
	})

}

func UpdateEmployee(c *gin.Context) {

	// Get id from url
	id := c.Param("id")

	// Get data from req body
	var reqBody struct {
		Employee_id          uint8 `gorm:"primary_key"`
		Employee_name        string
		Employee_age         uint8
		Employee_salary      float32
		Employee_Designation string
	}

	c.Bind(&reqBody)

	// Get employee
	var employee models.Employees
	initializer.DB.First(&employee, id)

	//Update user
	initializer.DB.Model(&employee).Updates(models.Employees{
		Employee_id:          reqBody.Employee_id,
		Employee_name:        reqBody.Employee_name,
		Employee_age:         reqBody.Employee_age,
		Employee_salary:      reqBody.Employee_salary,
		Employee_Designation: reqBody.Employee_Designation})

	c.JSON(http.StatusOK, gin.H{
		"Employee": employee,
	})
}

func DeleteEmployee(c *gin.Context) {

	// Get id from url
	id := c.Param("id")

	// Delete employee
	result := initializer.DB.Delete(&models.Employees{}, id)

	if result.Error != nil {
		c.Status(http.StatusUnprocessableEntity)
		return
	}

	c.Status(http.StatusOK)
}
