package main

import "fmt"

// ts
type employee struct {
	employeeID   string
	employeeName string
	phone        string
}

func main() {
	employeeList := []employee{}
	employee1 := employee{
		employeeID:   "101",
		employeeName: "AAA",
		phone:        "11111111",
	}
	employee2 := employee{
		employeeID:   "102",
		employeeName: "BBB",
		phone:        "222222",
	}

	employeeList = append(employeeList, employee1, employee2)

	fmt.Println("employee = ", employeeList)

}
