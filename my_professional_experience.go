package goEssentials

import "fmt"

func functionThree() {

	var name string = "Badar Ali" // normal decalation
	var age = 30                  // without dataType declartion
	country := "Austria"          // short declation

	isLearningGo := true
	jobTitle := "Backend engineer"
	experience := 4.5

	var (
		height = 185
		weight = 120
	)

	//new way of printing
	/*
			formated values are:
			%s - string
		    %d - integer
		    %f - float
		    %t - boolean
		    %v - any value (generic)

	*/
	fmt.Printf("This line is printed via printf function where my professional data is %s age %d country %s  weight %dkg and height %dcm is learning go %t and have job title as %s with %0.2f of experience", name, age, country, weight, height, isLearningGo, jobTitle, experience)

}
