package main

import (
	"fmt"
	"math"
	"reflect"
)

/*
a struct is a type which contains named fields
type Circle struct{
  x float64
  y float64
  z float64
}

type introduces a new type, followed by the name, the struct keywork and the
body, a curly braces block with field names and their corresponding types

note: like with functions, we can collapse fields with the same type
type Circle struct{
  x, y, z float64
}

we can create an instance of a struct
1) "declare + define"
var c Circle // like with other data types, this creates an instance of Circle
             // set to its zero value, for a struct this means each field is
             // set to its zero (0 int, 0.0 float64, "" string, nil pointer)

2) new() function
c := new(Circle) // allocates memory for all the fields, sets them to zero and
                 // returns a pointer to the memory, pointers to structs allow
                 // functions to modify their contents

3) initialize fields
c := Circle{x: 0, y: 0, r: 5}
c := Circle{0, 0, 5}
c := &Circle(0, 0, 5)

access struct fields using dot notation
fmt.Println(c.x, c.y, c.r)
c.x = 10
c.y = 5

e.g. function which computes a circle's area
func circleArea(c Circle) float64{
  return math.Pi * c.r * c.r
}

e.g. using a pointer to a circle
note: dereferencing a pointer seems to be done with the dot operator as well
func circlePtr(c Circle){
  return math.Pi * c.r * c.r
}

methods are functions linked to a struct type

// in between the func keyword and the parameter list, we've added a receiver
// which specifies the name and type of the object the method acts on, now the
// method can be called with dot notation
// note: the Golang runtime will automatically pass a pointer if the receiver
// says so, e.g. c.areaMethod() // where c is a struct
func (this *Circle) area() float64{
  return math.Pi * this.r * this.r
}

fields describe a HAS A relationship, structs can be embedded into other
structs to describe an IS A relationship, embedded types are also known as
anonymous fields

// person has a name
type Person struct{
  Name string
}
func (p *Person) Talk(){
  fmt.Println("hi my name is", p.name)
}

// android is a person
type Android2{
  Person
  Model string
}

now the person struct can be accessed usign the type name
a := new(Android)
a.Person.talk()

or person methods can be called directly on the android instance
a := new(Android) // a is a pointer
a.Talk()

intuitively, an android is a person, therefore an android can talk

notice that the rectangle's area method and the circle's share the same name
in the real world, relationships like these are common, go has a way of making
the relationship explicit through a type known as an interface

// instead of fields, an interface defines a method set
// a type implementing the interface must have all the specified methods
// interface types can be used as parameters in functions
type Shape interface{
  area() float64
}
*/

type Person struct {
	Name string
}

func (p *Person) Talk() {
	fmt.Println("hi my name is", p.Name)
}

type Circle struct {
	x float64
	y float64
	r float64
}

type Rectangle struct {
	x1 float64
	y1 float64
	x2 float64
	y2 float64
}

// this struct seems to indicate that an Android1 has a person
type Android1 struct {
	Person Person
	Model  string
}

type Android2 struct {
	Person // embedded type
	Model  string
}

func (r *Rectangle) area() float64 {
	l := distance(r.x1, r.y1, r.x1, r.y2)
	w := distance(r.x1, r.y1, r.x2, r.y1)
	return l * w
}

func distance(x1, y1, x2, y2 float64) float64 {
	dx := x2 - x1
	dy := y2 - y1
	return math.Sqrt(dx*dx + dy*dy)
}

func (this *Circle) area() float64 {
	return math.Pi * this.r * this.r
}

func circleArea(c *Circle) float64 {
	return math.Pi * c.r * c.r
}

func main() {
	// struct
	c := Circle{0, 0, 5}
	fmt.Println(reflect.TypeOf(c))
	fmt.Println(circleArea(&c))

	// pointer
	c2 := &Circle{0, 0, 5}
	fmt.Println(reflect.TypeOf(c2))
	fmt.Println(circleArea(c2))

	// method calls
	c3 := Circle{1, 1, 10}
	fmt.Println("struct method:", c3.area())

	r := Rectangle{0, 0, 1, 1}
	fmt.Println(r.area())

	// HAS A
	p := Person{"pythagroean"}
	p.Talk()

	// IS A
	// curly braces require commas (initialization)
	// parentheses need new line (import)
	a := Android2{
		Person{"jackson"},
		"A",
	}
	a.Talk()

	// fields are also inherited from embedded types
	fmt.Println(a.Name)
}
