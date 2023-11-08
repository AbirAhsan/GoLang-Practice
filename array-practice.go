package main

import "fmt"

func arrayPractice() {

	// declare fixed length array
	var studentList [6]int = [...]int{1, 10, 20, 30, 59, 593}

	fmt.Println(studentList)
	fmt.Println(len(studentList))                     // length
	fmt.Printf("Type is %T  \n", studentList)         // Type
	fmt.Println("3rd element value ", studentList[2]) // Index value
	//modify array elemet
	studentList[3] = 1
	fmt.Println(studentList)

	// Iterating over an Array
	for i, student := range studentList {
		fmt.Printf("index : %d, value %d\n", i, student)
	}

}
