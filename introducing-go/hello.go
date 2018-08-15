// every go program starts with a package declaration
// go programs can be executables or libraries, main declares an executable
package main

// fmt is a package which provides input and output formatting
// the package would possibly start with package fmt
import "fmt"

// main function declaration (program entry point)
// functions have input, output and statements to execute
// func <name> (<params>) <optional return type> {<body>}
//
// note that the opening curly brace must appear on the same line and that
// semicolons are inserted by the compiler
func main(){
	// access Println function of the fmt package
	fmt.Println("hello world")
}
