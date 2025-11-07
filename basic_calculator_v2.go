package goEssentials

//package main

import "fmt"

// TODO: Create these functions:

// 1. Add function - takes two integers, returns their sum
func add(a int, b int) int {
	return a + b
}

// 2. Multiply function - takes two integers, returns their product
func multiply(a int, b int) int {
	return a * b
}

// 3. Introduce function - takes name and age, returns formatted string
func introduce(name string, age int) string {
	data := fmt.Sprintln("hello ", name, " is your age is ", age)
	return data
}

func isAdult(age int) bool {
	if age >= 18 {
		return true
	}
	return false
}

func functionSix() {
	sum := add(15, 25)
	product := multiply(8, 7)
	introduction := introduce("Badar", 30)
	bmi := calculateBMI(120, 185)
	adultCheck := isAdult(30)

	fmt.Printf("Sum: %d, Product: %d\n", sum, product)
	fmt.Print(introduction) // Using Print since introduce already has newline
	fmt.Printf("BMI: %.2f\n", bmi)
	fmt.Printf("Is adult: %t\n", adultCheck)
}

func calculateBMI(weight int, height int) float64 {
	heightInMeters := float64(height) / 100.0
	return float64(weight) / (heightInMeters * heightInMeters)
}
