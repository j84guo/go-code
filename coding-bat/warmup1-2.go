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

/*
* utility
*/
func isTeen(i int) bool{
	return i >= 13 && i <= 19
}
