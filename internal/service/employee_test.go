package service

import (
	"ems/internal/dto"
	"ems/internal/repository"
	"github.com/stretchr/testify/suite"
	"testing"
)

func TestNewEmployeeServiceImpl(t *testing.T) {
	suite.Run(t, new(EmployeeServiceImplTestSuite))
}

type EmployeeServiceImplTestSuite struct {
	suite.Suite
	service *EmployeeServiceImpl
}

func (s *EmployeeServiceImplTestSuite) SetupTest() {
	s.service = NewEmployeeServiceImpl(repository.NewEmployeeInMemory())
}

func (s *EmployeeServiceImplTestSuite) TestCreate() {
	employee := &dto.EmployeeCreateUpdateRequest{Name: "Test", Salary: 1000, Position: "Developer"}
	expected := &dto.Employee{Name: "Test", Salary: 1000, Position: "Developer"}
	got, err := s.service.Create(employee)

	// Check if there is no error
	s.NoError(err)

	// Compare the expected and got
	s.Equal(expected.Name, got.Name)
	s.Equal(expected.Salary, got.Salary)
	s.Equal(expected.Position, got.Position)
}

func (s *EmployeeServiceImplTestSuite) TestGetNotFound() {
	_, err := s.service.Get(1)
	s.Error(err)
}

func (s *EmployeeServiceImplTestSuite) TestGetFound() {
	employeeDto := &dto.EmployeeCreateUpdateRequest{Name: "Test", Salary: 1000, Position: "Developer"}
	created, err := s.service.Create(employeeDto)
	s.NoError(err)

	got, err := s.service.Get(created.ID)
	s.NoError(err)
	s.Equal(created, got)
}
