package main

import "fmt"

func main() {
	fizzbuzz()
}

func fizzbuzz() {
	for i := 1; i <= 100; i++ {
		var three int = i % 3
		var five int = i % 5

		if (three == 0) && (five == 0) {
			fmt.Println(i, " fizzbuzz")
		} else if three == 0 {
			fmt.Println(i, " fizz")
		} else if five == 0 {
			fmt.Println(i, " buzz")
		}
	}
}
