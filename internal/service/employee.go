package service

import "ems/internal/dto"

type EmployeeService interface {
	Create(employee *dto.EmployeeCreateUpdateRequest) (*dto.Employee, error)
	Update(id int, employee *dto.EmployeeCreateUpdateRequest) (*dto.Employee, error)
	Delete(id int) error
	Get(id int) (*dto.Employee, error)
	List(pageNumber, size int) ([]*dto.Employee, error)
}
