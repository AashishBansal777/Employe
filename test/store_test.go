package test

import (
	"employe/pkg/store"
	"testing"
)

func TestCreateEmployee(t *testing.T) {
	store := store.NewEmployeeStore()
	emp := store.CreateEmployee("John Doe", "Developer", 60000)

	if emp.ID != 1 {
		t.Errorf("Expected ID to be 1, got %d", emp.ID)
	}
	if emp.Name != "John Doe" {
		t.Errorf("Expected Name to be John Doe, got %s", emp.Name)
	}
}

func TestGetEmployeeByID(t *testing.T) {
	store := store.NewEmployeeStore()
	emp := store.CreateEmployee("John Doe", "Developer", 60000)
	fetchedEmp, exists := store.GetEmployeeByID(emp.ID)

	if !exists {
		t.Errorf("Expected employee to exist")
	}
	if fetchedEmp.ID != emp.ID {
		t.Errorf("Expected ID to be %d, got %d", emp.ID, fetchedEmp.ID)
	}
}

func TestUpdateEmployee(t *testing.T) {
	store := store.NewEmployeeStore()
	emp := store.CreateEmployee("John Doe", "Developer", 60000)
	updatedEmp, exists := store.UpdateEmployee(emp.ID, "Jane Doe", "Manager", 80000)

	if !exists {
		t.Errorf("Expected employee to exist")
	}
	if updatedEmp.Name != "Jane Doe" {
		t.Errorf("Expected Name to be Jane Doe, got %s", updatedEmp.Name)
	}
}

func TestDeleteEmployee(t *testing.T) {
	store := store.NewEmployeeStore()
	emp := store.CreateEmployee("John Doe", "Developer", 60000)
	success := store.DeleteEmployee(emp.ID)

	if !success {
		t.Errorf("Expected deletion to be successful")
	}
	_, exists := store.GetEmployeeByID(emp.ID)
	if exists {
		t.Errorf("Expected employee to be deleted")
	}
}

func TestListEmployees(t *testing.T) {
	store := store.NewEmployeeStore()
	store.CreateEmployee("John Doe", "Developer", 60000)
	store.CreateEmployee("Jane Doe", "Manager", 80000)
	employees := store.ListEmployees()

	if len(employees) != 2 {
		t.Errorf("Expected 2 employees, got %d", len(employees))
	}
}
