package employee

import "fmt"

type employee struct {
	FirstName   string
	LastName    string
	TotalLeave  int
	LeavesTaken int
}

var employeeInstance *employee

func Init(
	firstName string,
	lastName string,
	totalLeave int,
	leavesTaken int) *employee {
	if employeeInstance == nil {
		employeeInstance = &employee{
			FirstName:   firstName,
			LastName:    lastName,
			TotalLeave:  totalLeave,
			LeavesTaken: leavesTaken}
	}
	return employeeInstance
}

func (e employee) LeavesRemaining() {
	fmt.Printf("%s %s has %d leaves remaining\n", e.FirstName, e.LastName, (e.TotalLeave - e.LeavesTaken))
}
