package goEssentials

import "fmt"

var name string = "Badar Ali"
var age = 30
var country = "Austria"

func functionEight() {

	basicDataFunction()
	fmt.Println("Professionalism: ", goingProfessionalFunction(4.5, "Backend engineer"))
}
func basicDataFunction() {
	fmt.Println("The name is", name, "age ", age, " country", country)
}

func goingProfessionalFunction(experience float32, jobTitle string) string {
	printable := fmt.Sprintln(" ", experience, " years of experience with job title as ", jobTitle)
	return printable
}
