package main

/*
Exercise: Student Grade Analyzer

Create a program that:
    Uses a struct to represent a Student (with Name and Grades)
    Stores multiple students in a slice
    Uses a map to track grade statistics
    Calculates and prints:
        Each student's average grade
        Class average
        Grade distribution (how many A, B, C, D, F)

Grade Scale:
    A: 90-100
    B: 80-89
    C: 70-79
    D: 60-69
    F: 0-59
*/

import "fmt"

type Student struct {
	Name   string
	Grades []int
}

func GradeAverage(randomSlice []int) int {
	sum := 0
	for _, valueData := range randomSlice {
		sum = valueData + sum
	}
	return (int)(sum / len(randomSlice))
}

func main() {

	//creating some students with their respected grades
	studentsSlice := []Student{
		{"Ali", []int{85, 92, 78}},
		{"Ahmad", []int{95, 98, 94}},
		{"Sadaqat", []int{65, 88, 81}},
		{"Bakr", []int{55, 70, 40}},
		{"Maroof", []int{30, 98, 68}},
		{"Dawood", []int{95, 98, 94}},
	}
	gradeStudentMap := make(map[string]int)
	studentMap := make(map[string]int)
	for _, valueData := range studentsSlice {
		studentMap[valueData.Name] = GradeAverage(valueData.Grades)
	}

	for mapKey, valueData := range studentMap {
		fmt.Println(" key: ", mapKey, " have grades average :", valueData)
	}

	fmt.Println("****************")

	classGradeSum := 0
	for _, valueData := range studentMap {
		classGradeSum = valueData + classGradeSum
	}
	fmt.Println("Class average is :", (int)(classGradeSum/len(studentMap)))

	fmt.Println("****************")

	for keyValueTemp, valueData := range studentMap {

		// var switchChecker = valueData
		fmt.Println(" key is: ", keyValueTemp, "value Data is :", valueData)
		switch {
		case valueData >= 90:
			gradeStudentMap["A"] = gradeStudentMap["A"] + 1
		case valueData >= 80:
			gradeStudentMap["B"] = gradeStudentMap["B"] + 1
		case valueData >= 70:
			gradeStudentMap["C"] = gradeStudentMap["C"] + 1
		case valueData >= 60:
			gradeStudentMap["D"] = gradeStudentMap["D"] + 1
		default:
			gradeStudentMap["F"] = gradeStudentMap["F"] + 1
		}
	}

	for keyValueTemp, valueData := range gradeStudentMap {

		fmt.Println(keyValueTemp, ": ", "Student(s) :", valueData)
	}
}
