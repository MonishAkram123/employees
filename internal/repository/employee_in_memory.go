package repository

import (
	"ems/internal/model"
	"ems/pkg/common"
	"errors"
	"sync"
	"time"
)

type EmployeeInMemoryRepository struct {
	employees       []*model.Employee
	employeeIndices map[int]int
	idGenerator     *common.IDGenerator

	mu *sync.Mutex
}

func NewEmployeeInMemory() *EmployeeInMemoryRepository {
	return &EmployeeInMemoryRepository{
		employees:       make([]*model.Employee, 0),
		employeeIndices: make(map[int]int),
		idGenerator:     common.NewIDGenerator(),
		mu:              &sync.Mutex{},
	}
}

func (repo *EmployeeInMemoryRepository) Get(id int) (*model.Employee, error) {
	idx, ok := repo.employeeIndices[id]
	if !ok {
		return nil, errors.New("employee not found")
	}
	return repo.employees[idx], nil
}

func (repo *EmployeeInMemoryRepository) Create(employee *model.Employee) (*model.Employee, error) {
	repo.mu.Lock()
	defer repo.mu.Unlock()

	employee.ID = repo.idGenerator.NextID()
	repo.employees = append(repo.employees, employee)
	repo.employeeIndices[employee.ID] = len(repo.employees) - 1
	return employee, nil
}

func (repo *EmployeeInMemoryRepository) Update(employee *model.Employee) (*model.Employee, error) {
	repo.employees[repo.employeeIndices[employee.ID]] = employee
	return employee, nil
}

func (repo *EmployeeInMemoryRepository) Delete(id int) error {
	repo.mu.Lock()
	defer repo.mu.Unlock()
	idx, ok := repo.employeeIndices[id]
	if !ok {
		return errors.New("employee not found")
	}

	repo.employees[idx].DeletedAt = common.GetPtr(time.Now())
	delete(repo.employeeIndices, id)
	return nil
}

func (repo *EmployeeInMemoryRepository) List(pageNumber, size int) ([]*model.Employee, error) {
	employees := repo.filterDeleted()
	start := (pageNumber - 1) * size
	end := start + size
	if start > len(employees) {
		return nil, nil
	}
	if end > len(employees) {
		end = len(employees)
	}
	return employees[start:end], nil
}

func (repo *EmployeeInMemoryRepository) filterDeleted() (employees []*model.Employee) {
	for _, employee := range repo.employees {
		if employee.DeletedAt == nil {
			employees = append(employees, employee)
		}
	}
	return
}
