package goEssential

import "fmt"

type Person struct {
	Name string
	Age  int
}

// Create a pointer using new()
// ptr := new(int)

// Function that takes a value (creates copy)
func addOneValue(n int) {
	n = n + 1
	fmt.Println("Inside addOneValue:", n)
}

// Function that takes a pointer (modifies original)
func addOnePointer(n *int) {
	*n = *n + 1
	fmt.Println("Inside addOnePointer:", *n)
}

func fourteenFunction() {

	someVaiable := 8
	fmt.Println("Value of someVariable: ", someVaiable)
	fmt.Println("Memory address of someVariable: ", &someVaiable)

	pointerVariable := &someVaiable
	fmt.Println("pointerVariable points towards: ", pointerVariable)
	fmt.Println("pointerVariable value: ", *pointerVariable)

	*pointerVariable = 100
	fmt.Println("Value after changing value of someVariable: ", someVaiable)
	fmt.Println("pointerVariable points towards: ", pointerVariable)
	fmt.Println("pointerVariable value: ", *pointerVariable)

	number := 5

	addOneValue(number)
	fmt.Println("After addOneValue:", number) // Still 5!

	addOnePointer(&number)
	fmt.Println("After addOnePointer:", number) // Now 6!

	// Regular struct
	person1 := Person{"Alice", 25}
	fmt.Println("Person1:", person1)

	// Pointer to struct
	person2 := &Person{"Bob", 30}
	fmt.Println("Person2:", person2)
	fmt.Println("Person2 Name:", person2.Name) // Go automatically dereferences!

	// Modify struct through pointer
	person2.Age = 31
	fmt.Println("Updated Person2:", person2)

}
