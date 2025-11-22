package main

/*
Zero Value Investigation

Write a program that declares variables of the following types without initializing them and then prints their values:

int
float64
string
bool
*int (a pointer to an int)
a struct with an int and a string field

Learning Goal: Solidify your knowledge of Go's zero values. You'll see that everything is initialized to a usable state (0, 0.0, "", false, nil).
*/

import "fmt"

type ShowStruct struct {
	someRandomByte byte
	someRandomRune rune
}

func main() {
	var age int
	var weight float64
	var isHuman bool
	var name string
	var randomPointer *int
	var randomStruct ShowStruct

	fmt.Printf("\n%T %T %T %T %T %T", age, weight, isHuman, name, randomPointer, randomStruct)
	fmt.Printf("\n%v %v %v %v %v %v\n", age, weight, isHuman, name, randomPointer, randomStruct)
}
