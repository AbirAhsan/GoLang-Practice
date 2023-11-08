package main

import "fmt"

func changeValueWithPointer(valuePonter *int) {
	*valuePonter = 6

}

// pointer with array
func modifyArray(a [5]*int) {
	*(a)[0] = 5
}

func pointerPractice() {

	var a int = 5
	fmt.Println("1st line a", a)

	changeValueWithPointer(&a)

	fmt.Println("2nd line a", a)

	//make function
	s := make([]int, 5)

	fmt.Println(s)

	arr := [5]int{1, 2, 3, 4, 5}
	fmt.Println(arr)
	// Create an array of pointers to integers
	var arrPointers [5]*int

	// Populate the array of pointers with the addresses of the elements in the original array
	for i := range arr {
		arrPointers[i] = &arr[i]
	}
	modifyArray(arrPointers)
	fmt.Println(arr)
}
