package main

import "fmt"

func main() {
	printHelloWorld()
	showVariable()
	constant()
	showZeroValues()
	showVariable1()
	showOperator()
	showIfConditions()
	a, b := returnParameters(5, 3)

	fmt.Println(a, b)
	value := 12
	passByValue(value)
	//
	passByreference(&value)
	fmt.Println("value change by reference", value)
	practiceFunction()
	arrayPractice()
	slicePractice()
}
