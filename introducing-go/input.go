package main

import "fmt"

func main() {
	fmt.Println("Enter a number: ")

	var input float64

	// pointer?
	fmt.Scanf("%f", &input)

	// type inference
	output := input * 2
	fmt.Println(output)
}
