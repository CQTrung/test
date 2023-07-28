package usecase

import "trunne-crud/entity"

type employeeUsecase struct {
	EmployeeRepository entity.EmployeeRepository
}

func NewEmployeeUsecase(er entity.EmployeeRepository) entity.EmployeeUsecase {
	return &employeeUsecase{
		EmployeeRepository: er,
	}
}

func (eu *employeeUsecase) FindAll() ([]entity.Employee, error) {
	employees, err := eu.EmployeeRepository.FindAll()
	if err != nil {
		return employees, err
	}
	return employees, nil
}
func (eu *employeeUsecase) FindById(id int) (entity.Employee, error) {
	employee, err := eu.EmployeeRepository.FindById(id)
	if err != nil {
		return employee, err
	}
	return employee, nil
}
func (eu *employeeUsecase) Create(employee entity.Employee) error {
	err := eu.EmployeeRepository.Save(employee)
	if err != nil {
		return err
	}
	return nil
}
func (eu *employeeUsecase) Update(employee entity.Employee) error {
	err := eu.EmployeeRepository.Save(employee)
	if err != nil {
		return err
	}
	return nil
}

func (eu *employeeUsecase) Delete(employeeID int) error {
	err := eu.EmployeeRepository.Delete(employeeID)
	if err != nil {
		return err
	}
	return nil
}
