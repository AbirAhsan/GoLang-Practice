package main

import "fmt"

func showVariable() {
	// with type
	var a int = 5
	fmt.Println(a)

	//without type
	b := 34.5

	fmt.Println(b)
}

func showVariable1() {
	var x int8 = 'a'
	var y uint8 = 255

	var f1 float32 = 4.8675775756565
	var f2 float64 = 4.8675775756565

	fmt.Println(x, y, f1, f2)
}
