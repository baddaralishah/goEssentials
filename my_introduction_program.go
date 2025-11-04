package goEssentials

import "fmt"

func someRandomName() {

	var name string = "Badar Ali" // normal decalation
	var age = 30                  // without dataType declartion
	country := "Austria"          // short declation

	var (
		height = 185
		weight = 120
	)

	fmt.Println(" Second line printing with name:", name, " age :", age, " living in:", country)
	fmt.Println(" Third line printing multiple variable declartion which are height:", height, " weight :", weight)

	//new way of printing
	/*
			formated values are:
			%s - string
		    %d - integer
		    %f - float
		    %t - boolean
		    %v - any value (generic)

	*/
	fmt.Printf("This line is printed via printf function where my data is %s age %d country %s  weight %dkg and height %dcm", name, age, country, weight, height)
}
