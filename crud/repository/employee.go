package repository

import (
	"trunne-crud/entity"

	"gorm.io/gorm"
)

type employeeRepository struct {
	Repository
}

func NewEmployeeRepository(database *gorm.DB) entity.EmployeeRepository {
	return &employeeRepository{
		Repository: Repository{
			database: database,
		},
	}
}

func (er *employeeRepository) FindAll() ([]entity.Employee, error) {
	var employees []entity.Employee
	err := er.database.Find(&employees).Error
	if err != nil {
		return employees, err
	}
	return employees, nil
}

func (er *employeeRepository) FindById(id int) (entity.Employee, error) {
	var employee entity.Employee
	err := er.database.First(&employee, id).Error
	if err != nil {
		return employee, err
	}
	return employee, nil
}

func (er *employeeRepository) Save(employee entity.Employee) error {
	err := er.database.Save(&employee).Error
	if err != nil {
		return err
	}
	return nil
}

func (er *employeeRepository) Delete(employeeID int) error {
	var employee entity.Employee
	err := er.database.First(&employee, employeeID).Error
	if err != nil {
		return err
	}
	err = er.database.Delete(&employee).Error
	if err != nil {
		return err
	}
	return nil
}
