package main

import "fmt"

func execute(x int, y int, operation func(x int, y int) int) {

	result := operation(x, y)

	fmt.Println(result)

}

// variadic function
func sumLikeVariadic(nums ...int) { //...int is for packing ( ellips)
	result := 0
	for _, num := range nums {
		result += num

		fmt.Println(result)
	}

	fmt.Println("result is", result)
}

func practiceFunction() {

	//anonymus function with return

	// func(x int, y int) int {
	// 	return x + y

	// }(5, 6)

	//anonymus function	and assing to a variable
	sum := func(x int, y int) int {
		return x + y
	}
	// result := sum(4, 9)
	// fmt.Println(result)

	// function as parameter
	execute(4, 5, sum)

	execute(4, 5, func(a int, b int) int {
		return a * b
	})

	// variadic function call
	value := []int{1, 2, 3, 4, 5}
	sumLikeVariadic(value...) // value... is for unpacking (sufficx ellips)

}
