package goEssentials

/*
Create a function that takes a number and:

Prints "Fizz" if divisible by 3
Prints "Buzz" if divisible by 5
Prints "FizzBuzz" if divisible by both

Otherwise prints the number
*/

import "fmt"

func customDividerWithIf(number int) string {
	if number%3 == 0 && number%5 == 0 {
		return "FizzBuzz"
	} else if number%5 == 0 {
		return "Buzz"
	} else if number%3 == 0 {
		return "Fizz"
	} else {
		returnNumber := fmt.Sprint("number is :", number)
		return returnNumber
	}
}

func customDividerWithSwitch(number int) string {
	switch {
	case number%3 == 0 && number%5 == 0:
		return "Fizz Buzz"
	case number%3 == 0:
		return "Fizz"
	case number%5 == 0:
		return "Buzz"
	default:
		returnNumber := fmt.Sprint("number is :", number)
		return returnNumber
	}
}

func functionFive() {

	counter := 1
	for counter < 23 {
		fmt.Println("looping counter ", counter, "with func customDividerWithIf function ", customDividerWithIf(counter))
		fmt.Println("looping counter ", counter, "with func customDividerWithSwitch function ", customDividerWithIf(counter))
		fmt.Println("*********************************")

		counter++
	}

}
