package main

import (
	"fmt"
  )

func main() {
var newEmployee = new(Employee)
newEmployee.Name = "Вася"
newEmployee.ShowDetails()
}


type Employee struct {
  Name string
  Age int
  Designation string
  Salary int
}


func (emp Employee) ShowDetails() {
  fmt.Println("User Name: ", emp.Name)
}

