package main

import (
	"errors"
	"fmt"
)

func divide(a, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("division by zero")
	}
	result := a / b
	return result, nil
}

type CustomError struct {
	Msg  string
	Code int
}

func (err CustomError) Error() string {
	return fmt.Sprintf("Error with code %d: %s", err.Code, err.Msg)
}

func SuperDivide(a, b int) (int, error) {

	if b == 0 {
		return 0, CustomError{Msg: "Super division by zero", Code: 1}
	}

	if a%b != 0 {
		return 0, CustomError{Msg: "What is error", Code: 2}
	}
	return a / b, nil
}

func errorPractice() {

	fmt.Println(divide(4, 0))

	res, err := divide(4, 1)

	if err != nil {
		fmt.Println("Error:", err.Error())
		return
	}
	fmt.Println("First1", res)

	res1, err1 := SuperDivide(5, 0)

	if err1 != nil {
		// Check if the error is of type CustomError
		if customErr, ok := err1.(CustomError); ok {
			// Print the error code
			fmt.Println("Error Code:", customErr.Code)
		}
		// Handle the error here
		fmt.Println("Error:", err)
	} else {
		// Use the result
		fmt.Println("Result: ", res1)
	}
	fmt.Println(res1)

}
