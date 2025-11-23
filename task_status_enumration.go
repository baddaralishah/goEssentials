package main

import "fmt"

type TaskStatus int

const (
	Pending TaskStatus = iota
	InProgress
	Completed
	Cancelled
)

func main() {

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
