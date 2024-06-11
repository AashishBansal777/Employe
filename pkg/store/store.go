package store

import (
	"sync"
)

type EmployeeStore struct {
	sync.Mutex
	employees map[int]Employee
	nextID    int
}

func NewEmployeeStore() *EmployeeStore {
	return &EmployeeStore{
		employees: make(map[int]Employee),
		nextID:    1,
	}
}

func (s *EmployeeStore) CreateEmployee(name, position string, salary float64) Employee {
	s.Lock()
	defer s.Unlock()

	emp := Employee{
		ID:       s.nextID,
		Name:     name,
		Position: position,
		Salary:   salary,
	}
	s.employees[s.nextID] = emp
	s.nextID++
	return emp
}

func (s *EmployeeStore) GetEmployeeByID(id int) (Employee, bool) {
	s.Lock()
	defer s.Unlock()

	emp, exists := s.employees[id]
	return emp, exists
}

func (s *EmployeeStore) UpdateEmployee(id int, name, position string, salary float64) (Employee, bool) {
	s.Lock()
	defer s.Unlock()

	emp, exists := s.employees[id]
	if !exists {
		return emp, false
	}

	emp.Name = name
	emp.Position = position
	emp.Salary = salary
	s.employees[id] = emp
	return emp, true
}

func (s *EmployeeStore) DeleteEmployee(id int) bool {
	s.Lock()
	defer s.Unlock()

	if _, exists := s.employees[id]; !exists {
		return false
	}
	delete(s.employees, id)
	return true
}

func (s *EmployeeStore) ListEmployees() []Employee {
	s.Lock()
	defer s.Unlock()

	employees := []Employee{}
	for _, emp := range s.employees {
		employees = append(employees, emp)
	}
	return employees
}
