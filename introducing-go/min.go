package main

import "fmt"

func main() {
	x := []int{
		14, 24, 64, 64,
		9, 14, 322, 21,
		51, 86, 33, 24,
		123, 43, 3, 11,
	}

	m := x[0]

	for _, v := range x {
		if v < m {
			m = v
		}
	}

	fmt.Println(m)
}
