package main

import(
  "fmt"
  "reflect"
)

/*
a function is an independent section of code which maps >=0 input to >=0 output

notes:
- arguments can be named differently than parameters (obviously)
- variables must be passed to functions (obviously)
- function calls are stacked
- RETURNED VARIABLES CAN HAVE NAMES SPECIFIED IN THE SIGNATURE
- MULIPLE VALUES CAN BE RETURNED

variadic functions may accept variable number of arguments

BASIS OF FUNCTIONAL PROGRAMMING:
closures are functions defined within another function
a function may call itself (recursion)
*/
func main(){
  reflectDemo()

  slice := []float64 {98, 93, 77, 82, 83}
  floatAverage(slice)

  // notImplemented()

  printX(10)

  f1()

  fmt.Println(returnName())

  // note that integers are initialized to 0
  var a int
  var b int

  a, b = two()
  fmt.Println(a, b)

  fmt.Println(add(1, 2, 3, 4))

  slice2 := []int{1, 2, 3, 4}
  fmt.Println(add(slice2...))

  // closure
  // doAdd is a local variable with type func(int, int) int, doAdd can also
  // access other global variables
  doAdd := func(x int, y int) int{
    return x + y
  }

  fmt.Println(doAdd(1, 11))

  closureDemo()

  var nextEven func() uint = makeEvenGenerator()
  fmt.Println(nextEven())
  fmt.Println(nextEven())
  fmt.Println(nextEven())

  fmt.Println(factorial(6))


}

func factorial(x uint) uint{
  if x == 0{
    return 1
  }
  
  return x * factorial(x - 1)
}

func makeEvenGenerator() func() uint{
  // cast constant to uint
  i := uint(0)

  return func() (r uint){
    r = i
    // the closure adds 2 to i
    // unlike normal local variables, the value of i persists between calls to nextInt()
    i += 2
    return
  }
}

/*
technically a function (nested in another) and the non-local variables it
accesses are known as a closure, so here x and inc are the closure

one way to use closures is writing a function which returns another function
*/
func closureDemo(){
  x := 0

  inc := func() int{
    x++
    return x
  }

  for i:=0; i<10; i++{
    fmt.Println(inc())
  }
}

// args is a slice of integers
// a variadic parameter must be the last in the signature
// note that a slice of ints can be passed by following the name with ...
func add(args ...int) (total int){
  total = 0

  fmt.Println(reflect.TypeOf(args))

  for _, v := range args{
    total += v
  }

  return
}

func two() (int, int){
  return 5, 6
}

func returnName() (res int){
  res = 90
  return
}

func f1(){
  fmt.Println("in f1...")
  f2()
  fmt.Println("...out f1")
}

func f2(){
  fmt.Println("in f2...")
  fmt.Println("...out f2")
}



func printX(x int){
  fmt.Println(x)
}

func reflectDemo(){
  // length 5 slice backed by a length 10 array
  var xs []float64 = make([]float64, 5, 10)
  fmt.Println(len(xs))
  fmt.Println(reflect.TypeOf(xs))
}

/*
functions start with the keyword func (easier parsing)
followed by the function name
parameters are included in parentheses as (name type, name type, ...)
followed by an optional return type

collectively this is known as the function signature
*/
func floatAverage(xs []float64) float64{
  var total float64 = 0.0

  for _, value := range xs{
    total += value
  }

  var average float64 = total / float64(len(xs))
  fmt.Println(average)

  return average
}

func notImplemented(){
  panic("notImplemented: function is unimplemented")
}
