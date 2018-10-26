package main

import (
	"errors"
	"fmt"
	"reflect"
)

func main() {
	//  *errors.errorString (pointer to errors.errorString)
	err := errors.New("my error message")
	fmt.Println(err, reflect.TypeOf(err))
}
