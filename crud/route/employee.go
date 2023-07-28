package route

import (
	"trunne-crud/controller"
	"trunne-crud/entity"

	"github.com/gin-gonic/gin"
)

func NewEmployeeRoute(group *gin.RouterGroup, employeeUsecase entity.EmployeeUsecase) {
	ec := controller.EmployeeController{EmployeeUsecase: employeeUsecase}
	group.GET("/employee/list", ec.GetAll)
	group.GET("/employee/detail", ec.GetByID)
	group.POST("/employee", ec.Create)
	group.PUT("/employee", ec.Update)
}
