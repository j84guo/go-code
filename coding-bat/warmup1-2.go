package main

import(
	"fmt"
	"strings"
	"reflect"
)

func main(){
	fmt.Println(
		hasTeen(1, 2, 3))

	fmt.Println(
		loneTeen(13, 19))

	fmt.Println(
		delDel("adelasd"))

	fmt.Println(
		mixStart("pixie"))

	fmt.Println(
		startOz("ozjackson"))

	fmt.Println(
		intMax(12, 2, 3))

	fmt.Println(
		close10(9, 11))

	fmt.Println(
		inRange(1, 10, 10))

	fmt.Println(
		in3050(39, 35))

    fmt.Println(
        stringE("jackson"))

    fmt.Println(
        stringE("jacksoneeee"))

    fmt.Println(
        stringE("ejaeckeson"))
}



func in3050(a, b int) bool{
	return inRange(a, 30, 40) && inRange(b, 30, 40) || inRange(a, 50, 40) && inRange(b, 50, 40)
}

func close10(a int, b int) int{
	da := abs(a - 10)
	db := abs(b - 10)

	if da < db{
		return a
	}else if db < da{
		return b
	}else{
		return 0
	}
}

func loneTeen(a int, b int) bool{
	var ab bool = isTeen(a)
	var bb bool = isTeen(b)

	return ab != bb
}

func hasTeen(a, b, c int) bool{
	sl := []int{a, b, c}

	for _, v := range sl{
		if isTeen(v){
			return true
		}
	}

	return false
}

func delDel(s string) string{
	if strings.Index(s, "del") == 1{
		return string(s[0]) + s[4:len(s)]
	}

	return s
}

func stringE(s string) bool{
    i := strings.Count(s, "e")
    return i >=1 && i <= 3
}

func mixStart(s string) bool{
	return strings.Index(s, "ix") == 1
}

func startOz(s string) string{
	var r string = ""

	fmt.Println("string: ", reflect.TypeOf(s))
	fmt.Println("string slice: ", reflect.TypeOf(s[0:2]))
	fmt.Println("string slice: ", reflect.TypeOf(s[0:1]))
	fmt.Println("byte: ", reflect.TypeOf(s[0]))

	if len(s) == 0{
		return s
	}

	if len(s) == 1{
		if(s == "0"){
			return s
		}else{
			return ""
		}
	}

	if s[0] == 'o'{
		r += string(s[0])
	}

	if s[1] == 'z'{
		r += string(s[1])
	}

	return r
}

func intMax(a, b, c int) int{
	sl := [3]int{a, b, c}
	max := sl[0]

	for _, v := range sl{
		if v > max{
			max = v
		}
	}

	return max
}

/*
* utility
*/
func isTeen(i int) bool{
	return i >= 13 && i <= 19
}

func abs(i int) int{
	if i < 0{
		i *= -1
	}

	return i
}

func inRange(i int, a int, b int) bool{
	if a > b{
		a, b = b, a
	}

	return i >= a && i <= b
}
