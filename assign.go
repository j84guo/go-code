package main

import "fmt"

func giveInt() (a, b int) {
	return 1, 2
}

func main() {
	a, b := giveInt()
	fmt.Println(a, b)

	a, b = giveInt()
	fmt.Println(":= needs at least one new variable")
}
