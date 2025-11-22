package goEssential

/*
Practice Exercise: Swap Function

Create a function that swaps the values of two integers using pointers.
Requirements:
    Function signature: func swap(a, b *int)
    Should swap the values that a and b point to
    Test it in main function
*/

import "fmt"

func swap(a, b *int) {
	temp := *a
	*a = *b
	*b = temp

	// same as
	// *a, *b = *b, *a
	//which is a parallel assignment feature in golang
}

func fifteenFunction() {

	x, y := 10, 20
	fmt.Printf("Before swap: x=%d, y=%d\n", x, y)

	swap(&x, &y)

	fmt.Printf("After swap: x=%d, y=%d\n", x, y)
	// Expected: x=20, y=10

}

/*

Key Learning Points:
    & operator - Gets the address of a variable (&x)
    * operator - Dereferences a pointer to get/set the value (*a)
    Pointer parameters - Allow functions to modify original variables
    No return needed - When using pointers, changes are reflected in the original variables

*/
