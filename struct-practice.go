package main

import (
	"fmt"
)

type Person struct {
	Name      string
	Age       int
	Gender    string
	isStudent bool
}

func (p Person) Print() { //Here (p Person) is receiver argument in function
	fmt.Printf("Name: %s, Age: %d, Gender: %s, isStudent: %t\n", p.Name, p.Age, p.Gender, p.isStudent)
}

func practiceStruct() {

	var p Person

	fmt.Println(p)

	// Assinging value
	p = Person{"Abdul", 25, "Male", true}
	fmt.Println(p)
	fmt.Printf("Name: %s, Age: %d, Gender: %s, isStudent: %t\n", p.Name, p.Age, p.Gender, p.isStudent)

	//Modifying fields
	p.Name = "Mr. Abdul"
	p.Print()

}
