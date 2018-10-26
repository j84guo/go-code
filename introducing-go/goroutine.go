package main

import (
	"fmt"
	"math/rand"
	"time"
)

func f(n int) {
	for i := 0; i < 10; i++ {
		fmt.Println(n, ":", i)
	}
}

func tf(n int) {
	for i := 0; i < 10; i++ {
		fmt.Println(n, ":", i)
		amt := time.Duration(rand.Intn(250))
		time.Sleep(time.Millisecond * amt)
	}
}

func main() {
	// twoDemo()
	// manyDemo()

	delayDemo()
}

/*
 * there are two goroutines, the twoDemo function itself and the invocation go
 * f(0), normally a function blocks the caller until it finishes, but a
 * goroutine creates a concurrent execution context, so the caller continues
 * immediately, Scanln is included so we can wait for the goroutine to finish
 */
func twoDemo() {
	go f(0)
	var input string
	fmt.Scanln(&input)
}

/*
 * goroutines are lightweight so we can easily create thousands of them
 */
func manyDemo() {
	for i := 0; i < 1000; i++ {
		go f(i)
	}
	var input string
	fmt.Scanln(&input)
}

/*
 * add random delay to each goroutine
 */
func delayDemo() {
	for i := 0; i < 10; i++ {
		go tf(i)
	}
	var input string
	fmt.Scanln(&input)
}
