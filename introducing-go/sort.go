package main

import (
	"flag"
	"fmt"
	"reflect"
	"sort"
)

type Person struct {
	Name string
	Age  int
}

/*
* define type ByName with underlying type slice of Person
*
* type Second First is a type definition, Second will inherit fields from First
* but not methods as they are bound to type First
 */
type ByName []Person

// slice does not have a Len method
func (ps ByName) Len() int {
	return len(ps)
}

// compare two elements in a slice
func (ps ByName) Less(i int, j int) bool {
	return ps[i].Name < ps[j].Name
}

// swap two elements in a slice
func (ps ByName) Swap(i int, j int) {
	ps[i], ps[j] = ps[j], ps[i]
}

type ByAge []Person

func (ps ByAge) Len() int {
	return len(ps)
}

func (ps ByAge) Less(i int, j int) bool {
	return ps[i].Age < ps[j].Age
}

func (ps ByAge) Swap(i int, j int) {
	ps[i], ps[j] = ps[j], ps[i]
}

func main() {
	kids := []Person{
		{"Jack", 90},
		{"Jill", 10},
	}

	/*
	* we define a type using out underlying slice to add methods Len, Less and
	* Swap
	*
	* function Sort takes a sort.Interface, which requires methods Len, Less
	* and Swap
	*
	* we then convert a []Person to ByName/ByAge, allowing it to be sorted
	 */
	sort.Sort(ByName(kids))
	fmt.Println(kids)

	sort.Sort(ByAge(kids))
	fmt.Println(kids)

	flag.Parse()
	args := flag.Args()
	fmt.Println(reflect.TypeOf(args), args)
	for _, s := range args {
		fmt.Println(s)
	}
}
