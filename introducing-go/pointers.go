package main

import (
	"fmt"
	"reflect"
)

/*
function arguments are normally passed by value

operators * and &:
- pointers are represented by * followwd by the type of stored value, e.g. *int
- * also de-references a pointer, and & returns the address of a variable
*/
func main() {
	x := 10

	notZero(x)
	fmt.Println(x)

	zero(&x)

	fmt.Println(x)
	fmt.Println(reflect.TypeOf(&x))

	// new is a built in function which returns a pointer
	// it takes a type as an argument, allocates enough memory for that type and
	// returns a pointer to the memory
	//
	// note that in go, objects are garbage collected, so dynamic allocation with
	// new does not need to be manually deleted
	xPtr := new(int)
	one(xPtr)

	// remember to de-reference xPtr
	fmt.Println(*xPtr)
}

func notZero(x int) {
	x = 0
}

func zero(xPtr *int) {
	*xPtr = 0 // de-reference x and assign it 0
}

func one(xPtr *int) {
	*xPtr = 1
}
