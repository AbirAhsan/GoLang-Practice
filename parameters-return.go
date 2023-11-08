package main

// func returnParameters(a int, b int) (int, int) {
// 	sum := a + b
// 	diff := a - b

// 	return sum, diff
// }

// another way
func returnParameters(a int, b int) (sum int, diff int) {
	sum = a + b
	diff = a - b

	return sum, diff // not mendatory return with variable name,
	// you can also write only return like this
	/*
			.....
			sum = a + b
			diff = a - b

			return
		}
	*/
}
