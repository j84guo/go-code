package main

import (
	"fmt"
	"reflect"
)

func main() {
	arr := [3]int{1, 2, 3}
	fmt.Println(reflect.TypeOf(arr))
	fmt.Println(arr.values)
}
