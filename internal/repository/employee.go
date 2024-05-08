package repository

import "ems/internal/model"

type EmployeeRepository interface {
	Create(employee *model.Employee) (*model.Employee, error)
	Update(employee *model.Employee) (*model.Employee, error)
	Delete(id int) error
	Get(id int) (*model.Employee, error)
	List(pageNumber, size int) ([]*model.Employee, error)
}
