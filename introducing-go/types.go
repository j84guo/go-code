package main

import "fmt"

// data types:
// 1) catgeorize a set of related values,
// 2) define properties and operations which can be performed on them
// 3) define how data is stored
//
// go is statically typed, meaning variables have a type which doesn't change
// static typing makes it easier to reason about what a program is doing and
// catches mistakes at compile time
//
// built in types:
//
// integer:
// numbers with no decimal component (computers use base 2 binary)
// when computers represent integers, all integers have a fixed size, eventually
// large numbers run out of space and wrap around
// uint8, uint16, uint32, uint64, int8, int16, int32, int64
// byte (uint8) and rune (int32) are aliases
// uint, int, uintptr are machine dependent
//
// floating point:
// decimal component
// floats are inexact
// like integers, floats have an exact width and precision increases with width
// in addition to numbers, there is also NaN and +/- infinity
// float32, float64 (single and double precision), complex64, complex128
//
// string:
// sequence of characters with a definite length
// many characters ads made up of one byte, but othres characters are multi-byte
// string literals are created using double quotes or ticks
// double quotes cannot contain newlines and allows escape sequences
//
// boolean:
// 1 bit integer type, true/false

func main() {
	fmt.Println("1 + 2 = ", 1+2)
	fmt.Println("1.0 + 2.0 = ", 1.0+2.0)

	// string length using global function len()
	fmt.Println(len("jackson"))

	// concatenation
	fmt.Println("hello " + "world")

	// character access
	fmt.Println("jackson"[0])

	fmt.Println(true || false)
}
