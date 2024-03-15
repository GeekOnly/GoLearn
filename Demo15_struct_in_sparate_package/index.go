package main

import (
	"demo15/employee"
)

func main() {
	e := employee.Employee{
		FirstName:   "Sanphet",
		LastName:    "Somjit",
		TotalLeave:  30,
		LeavesTaken: 20,
	}

	e.LeavesRemaining()
}
