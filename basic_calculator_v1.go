package goEssentials

//package main

import "fmt"

// TODO: Create these functions:

// 1. Add function - takes two integers, returns their sum
func addOne(a int, b int) int {
	return a + b
}

// 2. Multiply function - takes two integers, returns their product
func multiplyOne(a int, b int) int {
	return a * b
}

// 3. Introduce function - takes name and age, returns formatted string
func introduceOne(name string, age int) string {
	data := fmt.Sprintln("hello ", name, " is your age is ", age)
	return data
}

func functionSeven() {
	// Test your functions here
	sum := add(15, 25)
	product := multiply(8, 7)
	introduction := introduce("Badar", 30)

	fmt.Printf("Sum: %d, Product: %d\n", sum, product)
	fmt.Println(introduction)
}
