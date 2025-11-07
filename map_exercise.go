package goEssential

import "fmt"

/*
Practice Exercise:
Create a program that:
    Stores a list of numbers in a slice
    Uses a map to count how many times each number appears
    Prints the counts
*/

func functionTwelve() {

	friendsSlice := []int{1, 1, 1, 2, 2, 5, 7, 7, 6, 6, 6, 8, 9}
	friendsSlice = append(friendsSlice, 4, 4, 8, 8, 8, 8)

	//intilize an empty map
	friendsMap := make(map[int]int)

	for _, existsData := range friendsSlice {

		friendsMap[existsData]++
		// _, existsInMap := friendsMap[existsData]
		// if existsInMap {
		// 	friendsMap[existsData] = friendsMap[existsData] + 1
		// } else {
		// 	friendsMap[existsData] = 1
		// }
		// fmt.Println("Slice Key is :", keyValue, " Slice value is :", existsData, " in map key :", existsData, " comes :", friendsMap[existsData])
	}

	for mapFinalKey, mapFinalValue := range friendsMap {
		fmt.Println("Key is map is :", mapFinalKey, " and the value is :", mapFinalValue)
	}

}
