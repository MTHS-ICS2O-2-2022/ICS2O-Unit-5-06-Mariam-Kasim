// Created by: Mariam Kasim
// Created on: March 2023
//
// This program finds does multiplication

package main

import (
	"fmt"
)

func main() {
	var number1 int
	var number2 int
	var counter = 0
	var answer = 0

	// input
	fmt.Println("Let's do multiplication!")
	fmt.Println("Enter the first number:")
	fmt.Scanln(&number1)
	fmt.Println("Enter the second number: ")
	fmt.Scanln(&number2)

	// process
	for counter < number2 {
		answer += number1
		counter++
	}

	// output
	fmt.Println("")
	fmt.Println(number1, " x ", number2, " = ", answer)
	fmt.Println("\nDone.")
}
