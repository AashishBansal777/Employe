package test

import (
	"employe/pkg/store"
	"testing"
)

func TestCreateEmployee(t *testing.T) {
	store := store.NewEmployeeStore()
	emp := store.CreateEmployee("Aashish", "Developer", 60000)

	if emp.ID != 1 {
		t.Errorf("Expected ID to be 1, got %d", emp.ID)
	}
	if emp.Name != "Aashish" {
		t.Errorf("Expected Name to be Aashish, got %s", emp.Name)
	}
}

func TestGetEmployeeByID(t *testing.T) {
	store := store.NewEmployeeStore()
	emp := store.CreateEmployee("Aashish", "Developer", 60000)
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
	emp := store.CreateEmployee("Aashish", "Developer", 60000)
	updatedEmp, exists := store.UpdateEmployee(emp.ID, "Karan", "Manager", 80000)

	if !exists {
		t.Errorf("Expected employee to exist")
	}
	if updatedEmp.Name != "Karan" {
		t.Errorf("Expected Name to be Karan, got %s", updatedEmp.Name)
	}
}

func TestDeleteEmployee(t *testing.T) {
	store := store.NewEmployeeStore()
	emp := store.CreateEmployee("Aashish", "Developer", 60000)
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
	store.CreateEmployee("Aashish", "Developer", 60000)
	store.CreateEmployee("Karan", "Manager", 80000)
	employees := store.ListEmployees()

	if len(employees) != 2 {
		t.Errorf("Expected 2 employees, got %d", len(employees))
	}
}
