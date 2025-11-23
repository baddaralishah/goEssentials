package main

/*
Basic iota Usage

Create constants for the days of the week using iota, but start the week from 1 (so Sunday = 1, Monday = 2, etc.).
Requirements:

    Use iota to auto-increment values
    Start counting from 1 instead of 0
    Include all 7 days

*/

import "fmt"

const (
	_ = iota
	SUNDAY
	MONDAY
	TUESDAY
	WEDNESDAY
	THURSDAY
	FRIDAY
	SATURDAY
)

func main() {

	dayMap := map[int]string{
		SUNDAY:    "Sunday",
		MONDAY:    "Monday",
		TUESDAY:   "Tuesday",
		WEDNESDAY: "Wednesday",
		THURSDAY:  "Thursday",
		FRIDAY:    "Friday",
		SATURDAY:  "Saturday",
	}

	fmt.Println("=== Days Using Map Loop ===")
	for value, name := range dayMap {
		fmt.Printf("%s: %d\n", name, value)
	}

}
