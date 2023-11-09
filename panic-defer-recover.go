package main

import "fmt"

// func deferRecover() {
// 	fmt.Println("Defer function")

//		if r := recover(); r != nil {
//			fmt.Println("Recover from", r)
//		}
//	}
func practicePanic() {
	//defer deferRecover()
	fmt.Println("Before panic")

	panic("Something went to wrong")

	// fmt.Println("After panic")
}
