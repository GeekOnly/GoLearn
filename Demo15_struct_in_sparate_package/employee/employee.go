package employee

import "fmt"

type Employee struct {
	FirstName   string
	LastName    string
	TotalLeave  int
	LeavesTaken int
}

func (e Employee) LeavesRemaining() {
	fmt.Printf("%s %s has %d leaves remaining\n", e.FirstName, e.LastName, (e.TotalLeave - e.LeavesTaken))
}
