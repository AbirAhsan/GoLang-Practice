package main

import "fmt"

func practiceMap() {
	//creating map
	var a map[string]int
	fmt.Println(a)
	//initialization
	a = map[string]int{
		"apple":  10,
		"banana": 12,
		"cherry": 9,
	}
	fmt.Println(a)
	//accessing elements
	fmt.Println(a["banana"])

	//modifing elements
	a["banana"] = 20

	fmt.Println(a["banana"])

	//add new elements
	a["orange"] = 15
	fmt.Println(a)
	// Iterating over an map

	for k, v := range a {
		fmt.Printf("Key %s, value %d\n", k, v)
	}

	//deleting elements
	delete(a, "banana")

	fmt.Println(a)
}
