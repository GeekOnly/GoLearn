package employee

import "fmt"

type employee struct {
	FirstName   string
	LastName    string
	TotalLeave  int
	LeavesTaken int
}

func New(
	firstName string,
	lastName string,
	totalLeave int,
	leavesTaken int) employee {
	e := employee{
		FirstName:   firstName,
		LastName:    lastName,
		TotalLeave:  totalLeave,
		LeavesTaken: leavesTaken}
	return e
}

func (e employee) LeavesRemaining() {
	fmt.Printf("%s %s has %d leaves remaining\n", e.FirstName, e.LastName, (e.TotalLeave - e.LeavesTaken))
}
