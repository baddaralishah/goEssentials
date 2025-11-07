package goEssential

import "fmt"

func functionElf() {

	friends := map[string]int{
		"Bakr":   31,
		"Maroof": 33,
		"Ali":    29,
	}

	//new entry
	friends["Sadaqat"] = 29

	//Access the values
	fmt.Println("Ali's age is: ", friends["Ali"])

	//checking if key exists
	/*
		The working is the double assignment here in the check. age will autmoatically
		be consider as value to wards exists key. Here exists key is friends["Ahmad"]
	*/
	if age, exists := friends["Ahmad"]; exists {
		fmt.Println("Details age of Ahmad key is: ", age)
	} else {
		fmt.Println("Details cant be fetched ... [Key not available]")
	}

	// iterating through the map
	for loopingKey, age := range friends {
		fmt.Println("looping keys are: ", loopingKey, " with age: ", age)
	}

}
