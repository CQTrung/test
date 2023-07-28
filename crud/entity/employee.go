package entity

type Employee struct {
	Id   int    `gorm:"primary_key:auto_increment" json:"id"`
	Name string `gorm:"type:varchar(255)" json:"name"`
	City string `gorm:"type:varchar(255)" json:"city"`
}
type EmployeeRepository interface {
	FindAll() ([]Employee, error)
	FindById(id int) (Employee, error)
	Save(employee Employee) error
	Delete(employeeID int) error
}
type EmployeeUsecase interface {
	FindAll() ([]Employee, error)
	FindById(id int) (Employee, error)
	Create(employee Employee) error
	Update(employee Employee) error
	Delete(employeeID int) error
}
