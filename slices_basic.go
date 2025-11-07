package goEssential

import "fmt"

func functionTen() {

	//slice is considered to be a dynamic array
	// declaration
	firstSlice := []int{50, 51, 52}

	// changing the slice
	firstSlice = append(firstSlice, 53, 54, 55)

	fmt.Println("full slice: ", firstSlice)
	fmt.Println("first element of the slice: ", firstSlice[0])
	fmt.Println("first two elements of the slice: ", firstSlice[:2])
	fmt.Println("from the second index: ", firstSlice[2:])
	fmt.Println("length of the slice: ", len(firstSlice))
}
