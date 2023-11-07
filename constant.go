package main

import "fmt"

func constant() {
	//iota
	const (
		First = iota
		Second
		Third
	)

	fmt.Println(First, Second, Third)
}
