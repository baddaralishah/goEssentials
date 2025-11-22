package main

import (
	"fmt"
)

/*
Package-Level vs. Block-Level Scope

Create a small program with two files (e.g., block_scope_understanding.go and package_scope_understanding.go).

In package_scope_understanding.go, declare a package-level variable var Version string = "1.0".
In main.go, print this Version variable.
Now, inside the main() function in main.go, declare another variable Version := "2.0".
Print the variable again. What happens?

Learning Goal: Understand how package-level variables are shared across files in the same package and how block-level (local) variables shadow package-level ones.
*/

func solvingPackageVariableShadowingProblem() {
	fmt.Println("package scope variable: ", Name)
	fmt.Println("package scope variable: ", family)
}

func main() {

	fmt.Println("package scope variable: ", Name)
	fmt.Println("package scope variable: ", family)
	var Name string = "AL E ALI"
	var family string = "Naqi"

	fmt.Println("local scope variable: ", Name)
	fmt.Println("package scope variable: ", family)
	fmt.Println("****Shadowing issue solution******")
	solvingPackageVariableShadowingProblem()
	fmt.Println("----Agian accessing local scope variable-----")
	fmt.Println("local scope variable: ", Name)
	fmt.Println("package scope variable: ", family)

}
