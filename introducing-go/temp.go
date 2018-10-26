package main

import "fmt"

func main() {
	fmt.Println("Enter a temperature in farenheit: ")

	var f float64
	fmt.Scanf("%f", &f)

	var c float64
	c = (f - 32) * (5.0 / 9.0)

	fmt.Println(c)
}
