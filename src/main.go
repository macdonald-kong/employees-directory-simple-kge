package main

import (
    "net/http"

    "github.com/gin-gonic/gin"
)

// employee represents data about an employee.
type employee struct {
    ID		string `json:"id"`
    JobTitle	string `json:"jobTitle"`
    Name	string `json:"name"`
    Email	string `json:"email"`
}

// employees slice to seed employee data.
var employees = []employee{
	{ID: "1", Name: "Chris Darvill", JobTitle: "Solutions Engineer Manager", Email: "chris@kongexample.com"},
	{ID: "2", Name: "Marco Marquez", JobTitle: "Solutions Engineer Manager", Email: "marco@kongexample.com"},
	{ID: "3", Name: "Mark Sivill", JobTitle: "Solutions Engineer", Email: "mark@kongexample.com"},
	{ID: "4", Name: "Sven Walter", JobTitle: "Solutions Engineer", Email: "sven@kongexample.com"},
	{ID: "5", Name: "Steffen Miller", JobTitle: "Solutions Engineer", Email: "steffen@kongexample.com"},
	{ID: "6", Name: "Bruno Mandic", JobTitle: "Solutions Engineer", Email: "bruno@kongexample.com"},
	{ID: "7", Name: "Hans Wallentin", JobTitle: "Solutions Engineer", Email: "hans@kongexample.com"},	
	{ID: "8", Name: "Misiu Pajor", JobTitle: "Solutions Engineer", Email: "misiu@kongexample.com"},	
	{ID: "9", Name: "David MacDonald", JobTitle: "Solutions Engineer", Email: "david@kongexample.com"},	
	{ID: "10", Name: "Jerome Guillaume", JobTitle: "Solutions Engineer", Email: "jerome@kongexample.com"},	
	{ID: "11", Name: "Pierre-Alexandre Loriot", JobTitle: "Solutions Engineer", Email: "pierre-alexandre@kongexample.com"},	
	{ID: "12", Name: "Deepak Grewal", JobTitle: "Solutions Engineer", Email: "deepak@kongexample.com"},	
	{ID: "13", Name: "Amrita Gupta", JobTitle: "Solutions Engineer", Email: "amrita@kongexample.com"},
	{ID: "14", Name: "Andrew Klitovchenko", JobTitle: "Solutions Engineer", Email: "andy@kongexample.com"},	
}

func main() {
    router := gin.Default()
    router.GET("/api/employees", getEmployees)
    router.GET("/api/employees/:id", getEmployeeByID)
    router.POST("/api/employees", postEmployees)

    router.Run("0.0.0.0:8080")
}

// getEmployees responds with the list of all employees as JSON.
func getEmployees(c *gin.Context) {
    c.IndentedJSON(http.StatusOK, employees)
}

// getEmployeesByID locates the employee whose ID value matches the id
// parameter sent by the client, then returns that employee as a response.
func getEmployeeByID(c *gin.Context) {
    id := c.Param("id")

    // Loop over the list of employees, looking for
    // an employee whose ID value matches the parameter.
    for _, a := range employees {
        if a.ID == id {
            c.IndentedJSON(http.StatusOK, a)
            return
        }
    }
    c.IndentedJSON(http.StatusNotFound, gin.H{"message": "employee not found"})
}

// postEmployees adds an employee from JSON received in the request body.
func postEmployees(c *gin.Context) {
    var newEmployee employee

    // Call BindJSON to bind the received JSON to
    // newEmployee.
    if err := c.BindJSON(&newEmployee); err != nil {
        return
    }

    // Add the new employee to the slice.
    employees = append(employees, newEmployee)
    c.IndentedJSON(http.StatusCreated, newEmployee)
}
