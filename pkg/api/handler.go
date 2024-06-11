package api

import (
	"employe/pkg/store"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

var empStore = store.NewEmployeeStore()

func NewRouter() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/employees", listEmployeesHandler).Methods("GET")
	router.HandleFunc("/employees", createEmployeeHandler).Methods("POST")
	router.HandleFunc("/employees/{id:[0-9]+}", getEmployeeHandler).Methods("GET")
	router.HandleFunc("/employees/{id:[0-9]+}", updateEmployeeHandler).Methods("PUT")
	router.HandleFunc("/employees/{id:[0-9]+}", deleteEmployeeHandler).Methods("DELETE")
	return router
}

func createEmployeeHandler(w http.ResponseWriter, r *http.Request) {
	var emp store.Employee
	if err := json.NewDecoder(r.Body).Decode(&emp); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	newEmp := empStore.CreateEmployee(emp.Name, emp.Position, emp.Salary)
	json.NewEncoder(w).Encode(newEmp)
}

func getEmployeeHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])

	emp, exists := empStore.GetEmployeeByID(id)
	if !exists {
		http.Error(w, "Employee not found", http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(emp)
}

func updateEmployeeHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])

	var emp store.Employee
	if err := json.NewDecoder(r.Body).Decode(&emp); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	updatedEmp, exists := empStore.UpdateEmployee(id, emp.Name, emp.Position, emp.Salary)
	if !exists {
		http.Error(w, "Employee not found", http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(updatedEmp)
}

func deleteEmployeeHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])

	if !empStore.DeleteEmployee(id) {
		http.Error(w, "Employee not found", http.StatusNotFound)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

func listEmployeesHandler(w http.ResponseWriter, r *http.Request) {
	page, _ := strconv.Atoi(r.URL.Query().Get("page"))
	if page < 1 {
		page = 1
	}

	pageSize, _ := strconv.Atoi(r.URL.Query().Get("pageSize"))
	if pageSize < 1 {
		pageSize = 10
	}

	employees := empStore.ListEmployees()
	paginatedEmployees := paginateEmployees(employees, page, pageSize)
	json.NewEncoder(w).Encode(paginatedEmployees)
}
