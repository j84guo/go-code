package main

import "fmt"

var globalHello string = "hello world (global)"

// a variable is a storage location with a specific type and name
func main(){
  // the string literal is assigned to a variable
  // variables are declared using the var keyword, followed by name and type
  var hello string = "hello world"
  fmt.Println(hello)

  var sum int
  sum = 10 + 10
  fmt.Println(sum)

  // variables can be assigned new values during their lifetime
  hello = "hello jackson"
  fmt.Println(hello)

  hello += " how are you"
  fmt.Println(hello)

  // type inference can be done using :=
  bye := "bye jackson"
  fmt.Println(hello == bye)

  // or var (excluding the type)
  var num = 3.14 // system probably uses a default floating point type
  fmt.Println(num)

  // variables names may contain letters, numbers or underscores
  fmt.Println(globalHello)

  // calling a function
  sayHello()

  // constants cannot be re-assigned
  const pi float64 = 3.14
  fmt.Println("the constant pi is ", pi)

  // compiler error
  // pi = 6.0

  // multiple variables can be defined at the same time
  var (
    a int = 5
    b int = 10
    c int = 15
  )

  // note that variables cannot be created and unused
  fmt.Println(a, b, c)
}

// the scope of a variable refers to the location in which it can be used
// go uses block scope, meaning variables exist in the current and nested {}
func sayHello(){
  fmt.Println(globalHello)
}
