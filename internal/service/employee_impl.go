package service

import (
	"ems/internal/dto"
	"ems/internal/model"
	"ems/internal/repository"
	"github.com/jinzhu/copier"
)

type EmployeeServiceImpl struct {
	repository repository.EmployeeRepository
}

func NewEmployeeServiceImpl(repository repository.EmployeeRepository) *EmployeeServiceImpl {
	return &EmployeeServiceImpl{repository: repository}
}

func (svc *EmployeeServiceImpl) Create(request *dto.EmployeeCreateUpdateRequest) (_ *dto.Employee, err error) {
	employee := &model.Employee{}
	if err = copier.Copy(employee, request); err != nil {
		return nil, err
	}

	response := &dto.Employee{}
	employee, err = svc.repository.Create(employee)
	if err != nil {
		return nil, err
	}

	err = copier.Copy(response, employee)
	return response, err
}

func (svc *EmployeeServiceImpl) Update(id int, request *dto.EmployeeCreateUpdateRequest) (_ *dto.Employee, err error) {
	employee, err := svc.repository.Get(id)
	if err != nil {
		return nil, err
	}

	if err = copier.Copy(employee, request); err != nil {
		return nil, err
	}
	employee.ID = id
	employee, err = svc.repository.Update(employee)
	if err != nil {
		return nil, err
	}

	response := &dto.Employee{}
	err = copier.Copy(response, employee)
	return response, err
}

func (svc *EmployeeServiceImpl) Delete(id int) error {
	return svc.repository.Delete(id)
}

func (svc *EmployeeServiceImpl) Get(id int) (employee *dto.Employee, err error) {
	employeeModel, err := svc.repository.Get(id)
	if err != nil {
		return nil, err
	}

	employee = &dto.Employee{}
	if err = copier.Copy(employee, employeeModel); err != nil {
		return nil, err
	}
	return employee, nil
}

func (svc *EmployeeServiceImpl) List(pageNumber, size int) (employees []*dto.Employee, err error) {
	employeeModels, err := svc.repository.List(pageNumber, size)
	if err != nil {
		return nil, err
	}

	employees = make([]*dto.Employee, 0)
	for _, employeeModel := range employeeModels {
		employee := &dto.Employee{}
		if err = copier.Copy(employee, employeeModel); err != nil {
			return nil, err
		}
		employees = append(employees, employee)
	}
	return employees, nil
}
