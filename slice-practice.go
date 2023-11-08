package main

import "fmt"

func slicePractice() {
	mySlice := []int{1, 2, 3, 4}
	fmt.Println(mySlice)

	// usong make
	mySlice2 := make([]int, 4)

	fmt.Println(mySlice2)

	mySlice2[3] = 40

	fmt.Println(mySlice2)

	//appending element
	mySlice2 = append(mySlice2, 20, 78, 30, 209, 43)
	fmt.Println("Slice 2", mySlice2)
	newSlice := mySlice2[4:7]
	fmt.Println("New slice", newSlice)
}
