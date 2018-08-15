package main

import(
  "fmt"
)

func main(){
  fmt.Println(
    sleepIn(true, true))

  fmt.Println(
    monkeyTrouble(false, false))

  fmt.Println(
    diff21(20))

  fmt.Println(
    parrotTrouble(true, 23))

  fmt.Println(
    makes10(10, 1))

  fmt.Println(
    nearOneTwoHundred(110))

  fmt.Println(
    posNeg(1, -1, true))
}

func sleepIn(weekday bool, vacation bool) bool{
  return vacation || !weekday
}

func monkeyTrouble(a bool, b bool) bool{
  return a && b || !a && !b
}

func sumDouble(a int, b int) int{
  sum := a + b

  if a == b{
    return 2 * sum
  }

  return sum
}

func diff21(n int) int{
  diff := abs(n - 21)

  if n > 21{
    return 2 * diff
  }

  return diff
}

func parrotTrouble(talk bool, hour uint) bool{
  if !talk{
    return false
  }

  return hour < 7 || hour > 20
}

func makes10(a int, b int) bool{
  if a == 10 || b == 10{
    return true
  }

  return a + b == 10
}

func nearOneTwoHundred(n int) bool{
  return abs(n - 100) <= 10 || abs(n - 200) <= 10
}

func posNeg(a int, b int, n bool) bool{
  an := a < 0
  bn := b < 0

  return n && an && bn || !n && an && !bn || !n && !an || bn
}

/*
* utilities
*/
func abs(n int) int{
  if n < 0{
    return n * -1
  }

  return n
}
