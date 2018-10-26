package main

import "fmt"

/*
slice is a segment of an array
slices are indexable and have length
unlike arrays, slice length is allowed to change

in addition to make() and index operations, go provides two built in functions
append() and copy()
*/
func main() {
	sliceDemo()
	makeDemo()
	boundsDemo()
	appendDemo()
	copyDemo()
}

func appendDemo() {
	// append adds items to the end of a slice, if there is no room in the
	// underlying array, a new one is created (resizable array), the new slice is
	// returned
	// note slices can be initialized
	slice1 := []int{1, 2, 3}
	slice2 := append(slice1, 4, 5)
	fmt.Println(slice1, slice2)
}

func copyDemo() {
	// copies entries from src into dst, stopping at the end of dst if reached
	slice1 := []int{2, 4, 6}
	slice2 := []int{12, 14, 16, 18}
	copy(slice1, slice2)

	// the last element of slice2 is untouched as it is longer than slice1
	fmt.Println(slice1, slice2)
}

func sliceDemo() {
	// notice a slice declaration is like an array but without a specified length
	// this creates a slice of length 0
	var x []float64
	fmt.Println(x)
}

func makeDemo() {
	// the built in make function creates a slice of length 5 associated with an
	// underlying array of length 5, all slices are associated with an array
	y := make([]float64, 5)

	// slice of length 5 associated with an array of length 10, slices cannoy be
	// larger than their array, but may be smaller
	z := make([]string, 5, 10)

	fmt.Println(y, z)
}

func boundsDemo() {
	arr := [5]float64{
		1, 2, 3, 4, 5,
	}

	// low is inclusive, high is exclusive
	x := arr[0:5]
	fmt.Println(x)
}
