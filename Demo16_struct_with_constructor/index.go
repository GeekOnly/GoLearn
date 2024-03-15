package main

import "demo16/employee"

func main() {
	e := employee.New(
		"Sanphet",
		"Somjit",
		30,
		20,
	)

	e.LeavesRemaining()
}
