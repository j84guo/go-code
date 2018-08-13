package main

import "fmt"

func main(){
  simpleDemo()
  accessDemo()
  rangeDemo()
  ignoreDemo()
  bracesDemo()
}

func simpleDemo(){
  // note that length is part of the array type name
  var arr[5] int
  arr[4] = 100
  fmt.Println(arr, arr[4])
}

func accessDemo(){
  var arr[5] float64
  arr[0] = 98
  arr[1] = 93
  arr[2] = 71
  arr[3] = 28
  arr[4] = 0

  var total float64 = 0
  for i:=0; i<len(arr); i++{
    total += arr[i]
  }

  // type conversion
  fmt.Println(total / float64(len(arr)))
}

func rangeDemo(){
  var arr[5] float64

  // the range keyword causes iteration over the array, i is the index and value
  // is the value in the array at i
  // note that _ is used for indicate an unused variable, otherwise the compiler
  // will give an unused variable error
  for i, value := range arr{
    fmt.Println(i, value)
  }
}

func ignoreDemo(){
  var arr[5] string

  for _, value := range arr{
    fmt.Println(value)
  }
}

func bracesDemo(){
  // array types can be inferred and initialized
  // note the trailing comma is required for multiline array initialization
  arr := [5] float64 {
    10,
    20,
    30,
    40,
    50,
  }

  fmt.Println(arr)
}
