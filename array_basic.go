package goEssential

import "fmt"

func functionNine() {

	// array declaration
	var firstArray [3]int
	//array assignment
	firstArray[0] = 15
	firstArray[1] = 20
	firstArray[2] = 25

	//short declaration and assignment
	secondArray := [len(firstArray)]int{30, 35, 40}

	fmt.Println("first array :", firstArray)
	fmt.Println("Second array :", secondArray)

}
