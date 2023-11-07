package main

import "fmt"

func showIfConditions() {

	var a int = 10

	if a > 10 {
		fmt.Println("False")
	} else if a == 10 {
		fmt.Println("true")
	} else {
		fmt.Println("False")
	}

	if valueDeclaration(a) {
		fmt.Println(a, "is positive")
	}

}

func valueDeclaration(value int) bool {
	return value > 0
}
