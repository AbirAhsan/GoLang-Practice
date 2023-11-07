package main

import "fmt"

func showOperator() {
	var a int = 4
	var b int = 6

	var c int = a + b
	var d int = a - b
	var e int = b / a
	var f int = a % b

	fmt.Println("sum =", c, "\ndifference=", d, "\ndivision=", e, "\nfraction=", f)
}
