package main

import "fmt"

func passByValue(value int) {

	fmt.Println("x value =", value)
}

func passByreference(value *int) {
	*value = 20
	//fmt.Println("y reference =", *value)
}
