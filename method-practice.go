package main

import "fmt"

type Dog struct {
	Name  string
	Age   uint
	Color string
}

func (d Dog) Bark() {
	fmt.Printf("%s says : woof!\n", d.Name)
}

func (d *Dog) changeDogName(name string) {
	d.Name = name
}

func methodPractice() {
	d := Dog{Name: "Rocy", Age: 3, Color: "White"}
	d.changeDogName("Booby")
	d.Bark()
}
