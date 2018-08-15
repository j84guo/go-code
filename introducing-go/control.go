package main

import "fmt"

// go has one looping keywork, for, which can be used in different ways (like
// traditional for and while) note parentheses are generally optional
// go also supports reasonably traditional if and switch statements, note that
// switch statements don't seem to require a break to prevent "fall through"
func main(){
  i := 1

  // while style
  for i <= 10 {
    fmt.Println(i)
    i++
  }

  // for style
  for j := 10; j >= 1; j-- {
    fmt.Println(j)
  }

  // if statements don't require parentheses
  for i = 1; i < 11; i++ {
    if i % 2 == 0 {
      fmt.Println(i, " even")
    } else {
      fmt.Println(i, " odd")
    }
  }

  // switch statement
  switch("jackson"){
  case "a":
    fmt.Println("a")
  case "jackson":
    fmt.Println("jackson")
  default:
    fmt.Println("default")
  }
}
