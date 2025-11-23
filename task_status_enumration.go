package goEssential

/*
Task Status Enumeration
Create a set of status constants for a task management system:
    Pending should have value 0
    InProgress, Completed, Cancelled should auto-increment using iota
Requirements:
    Explicitly set Pending = 0
    Use iota for the remaining statuses
    Make it a typed constant for better type safety
*/
import "fmt"

type TaskStatus int

const (
	Pending TaskStatus = iota
	InProgress
	Completed
	Cancelled
)

func functionNinteen() {

	statusMap := map[TaskStatus]string{
		Pending:    "PENDING",
		InProgress: "INPROGRESS",
		Completed:  "COMPLETED",
		Cancelled:  "CANCELLED",
	}

	for indexValue, statusCurrent := range statusMap {
		fmt.Println(" status ", statusCurrent, " with iota value: ", indexValue)
	}
}
