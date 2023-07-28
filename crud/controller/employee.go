package controller

import (
	"strconv"
	"trunne-crud/entity"

	"github.com/gin-gonic/gin"
)

type EmployeeController struct {
	EmployeeUsecase entity.EmployeeUsecase
}

func (ec *EmployeeController) Create(c *gin.Context) {
	var employee entity.Employee
	c.BindJSON(&employee)
	ec.EmployeeUsecase.Create(employee)
	c.JSON(200, employee)
}

func (ec *EmployeeController) GetAll(c *gin.Context) {
	employees, err := ec.EmployeeUsecase.FindAll()
	if err != nil {
		c.JSON(500, err)
		return
	}
	c.JSON(200, employees)
}

func (ec *EmployeeController) GetByID(c *gin.Context) {
	var id, _ = strconv.Atoi(c.Query("id"))
	employee, err := ec.EmployeeUsecase.FindById(id)
	if err != nil {
		c.JSON(500, err)
		return
	}
	c.JSON(200, employee)
}

func (ec *EmployeeController) Update(c *gin.Context) {
	var employee entity.Employee
	c.BindJSON(&employee)
	ec.EmployeeUsecase.Update(employee)
	c.JSON(200, employee)
}
