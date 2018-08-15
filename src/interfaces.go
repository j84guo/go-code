package main

import(
  "fmt"
  "math"
)

/*
an interface is a type which defines a method set, a list of methods a struct
must have to implement the interface

interface types can be used as parameters in functions
note that NO KEYWORD is need to implement an interface, simply having the
appropriate methods is enough

interfaces can also be used as fields

when building a program, we often don't know what types are needed at the start
Go generally focuses on the behaviour of a program rather than on types
(procedural vs object-oriented)

create some small structs which do what you want, add methods as needed and
useful interfaces will start to emerge as the program develops

interfaces are useful as software grows larger, to hide implementation details
which makes it easier to reason about software components in isolation

Go also has a mechanism, packages, for combining interfaces, types, variables
and functions into a common namespace
*/

type Shape interface{
  area() float64
}

type MultiShape struct{
  shapes []Shape
}

type Circle struct{
  x float64
  y float64
  r float64
}

type Rectangle struct{
  x1 float64
  y1 float64
  x2 float64
  y2 float64
}

// pointer receiver
func (ms *MultiShape) area() float64{
  var area float64 // initialized to zero value by default

  for _, s := range ms.shapes{
    area += s.area()
  }

  return area
}

// copy receiver
func (r Rectangle) area() float64{
  l := distance(r.x1, r.y1, r.x1, r.y2)
  w := distance(r.x1, r.y1, r.x2, r.y1)
  return l * w
}

func distance(x1, y1, x2, y2 float64) float64{
  dx := x2 - x1
  dy := y2 - y1
  return math.Sqrt(dx*dx + dy*dy)
}

// copy receiver
func (this Circle) area() float64{
  return math.Pi * this.r * this.r
}

// this function works, but we must add a new slice parameters for each type of
// shape we want to handle
func totalAreaDumb(cs []Circle, rs []Rectangle) float64{
  var total float64 = 0

  for _, c := range cs{
    total += c.area()
  }

  for _, r := range rs{
    total += r.area()
  }

  return total
}

// use an interface as a variadic parameter instead
// all total area knows is that the items in the slice have a shape method, so
// the function would not be able to access fields or other methods
func totalArea(shapes ...Shape) float64{
  total := float64(0)

  for _, sh := range shapes{
    total += sh.area()
  }

  return total
}

func main(){
  c := Circle{0, 0, 9}

  r := Rectangle{0, 0, 5, 5}

  // note that the shapes are pased in as pointers
  fmt.Println(totalArea(&c, &r))

  // interfaces as fields
  gms := MultiShape{
    shapes: []Shape{
      Circle{0, 0, 5},
      Rectangle{0, 0, 1, 1},
    },
  }

  fmt.Println(gms)
  fmt.Println(gms.area())
}
