package main

import(
  "fmt"
  "strings"
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

  fmt.Println(
	notString("not a string"))

  fmt.Println(
    missingChar("jackson", 1))

  fmt.Println(
	frontBack("ja"))

  fmt.Println(
    backAround("jackson"))

  fmt.Println(
    front22("jackson"))

  fmt.Println(
    startHi("hi world"))
  fmt.Println(
    startHi("hello world"))

  fmt.Println(
    icyHot(10, 1100))

  fmt.Println(
    in1020(10, 13))
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

func notString(s string) string{
  if strings.HasPrefix(s, "not"){
    return s
  }

  return "not " + s
}

func missingChar(s string, i int) string{
  if i < 0 || i > len(s){
    panic("missingChar: invalid index")
  }

  return s[0:i] + s[i+1:len(s)]
}

func frontBack(s string) string{
  if len(s) < 2{
    return s
  }

  return string(s[len(s)-1]) + s[1:len(s)-1] + string(s[0])
}

func front3(s string) string{
	var base string
	if len(s) < 3{
		base = s
	}else{
		base = s[0:2]
	}

	return base + base + base
}

func backAround(s string) string{
  if len(s) == 0{
    return ""
  }

  var c byte = s[len(s) - 1]
  return string(c) + s + string(c)
}

func or35(n uint) bool{
  return n % 3 == 0 || n % 5 == 0
}

func front22(s string) string{
  if len(s) < 2{
    return s + s + s
  }

  a := s[0:2]
  return a + s + a
}

func startHi(s string) bool{
  return strings.HasPrefix(s, "hi")
}

func icyHot(a int, b int) bool{
  return a < 0 && b > 100 || b < 0 && a > 100
}

func in1020(a int, b int) bool{
  return a >= 10 && a <= 20 || b >= 10 && b <= 20
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
