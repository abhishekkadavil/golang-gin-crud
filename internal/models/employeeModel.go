package models

type Employees struct {
	Employee_id          uint8 `gorm:"primary_key"`
	Employee_name        string
	Employee_age         uint8
	Employee_salary      float32
	Employee_Designation string
}
