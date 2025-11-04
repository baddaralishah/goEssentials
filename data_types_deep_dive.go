package goEssentials

import "fmt"

func someRandomName() {

	/*
		Variable complete declation
	*/

	var integer int = 30
	var floatingPoint float64 = 4.5
	var boolean bool = true
	var text string = "Hello Go!"

	// Zero Values (default values)
	var defaultInt int       // 0
	var defaultFloat float64 // 0.0
	var defaultBool bool     // false
	var defaultString string // ""

	fmt.Printf("Integer: %d, Float: %f, Boolean: %t, String: %s\n",
		integer, floatingPoint, boolean, text)
	fmt.Printf("Defaults - Int: %d, Float: %f, Bool: %t, String: '%s'\n",
		defaultInt, defaultFloat, defaultBool, defaultString)
}
