package main

import(
    "fmt"
    "reflect"
)

func main(){
  sp := []int{
    1,
    2,
    3,
    4,
    5,
  }

  fmt.Println(sumSplice(sp))

  var(
    i int
    b bool
  )

  i, b = halfEvenOdd(10)
  fmt.Println(i, b)

  fmt.Println(max(1, 2, 3))
  fmt.Println(max(10, 24, 23))

  odds := makeOddGenerator()
  fmt.Println(odds())
  fmt.Println(odds())
  fmt.Println(odds())

  fmt.Println(simpleFib(10))
  fmt.Println(fib(10))

  deferPanicRecover()

  var n int = 10
  address(&n)

  setTwo(&n)
  fmt.Println(n)

  var f *float64 = new(float64)
  fmt.Println(*f)

  x := 1.5
  square(&x)
  fmt.Println(x)

  first := 1
  second := 2
  swap(&first, &second)
  fmt.Println(first, second)
}

func swap(a *int, b *int){
  tmp := *a
  *a = *b
  *b = tmp
}

func square(x *float64){
  *x = *x * *x
}

func setTwo(n *int){
  *n = 2
}

func address(n *int){
  fmt.Println(n)
}

func deferPanicRecover(){
  // closure
  rec := func(){
    s := recover()
    fmt.Println(s)
  }

  defer rec()

  panic("panic")
}

func fib(n uint) uint{
  cache := make(map[uint]uint)

  return fibRec(n, cache)
}

func fibRec(n uint, cache map[uint]uint) uint{
  if n <= 1{
    return n
  }

  if v, ok := cache[n]; ok == true{
    return v
  }else{
    return fibRec(n-1, cache) + fibRec(n-2, cache)
  }
}

func simpleFib(n uint) uint{
  if n <= 1{
    return n
  }

  return simpleFib(n-1) + simpleFib(n-2)
}

func makeOddGenerator() func() int{
  var x int = -1

  return func () int{
    x += 2
    return x
  }
}

func max(args ...int) int{
  fmt.Println(args)
  fmt.Println(reflect.TypeOf(args))
  fmt.Println("len:", len(args))

  if len(args) == 0{
    panic("max: must have at least 1 argument")
  }

  m := args[0]

  for _, v := range args{
    if v > m{
      m = v
    }
  }

  return m
}

func sumSplice(sp []int) (total int){
  total = 0

  for _, v := range sp{
    total += v
  }

  return
}

func halfEvenOdd(n int) (int, bool){
  h := n / 2
  b := n % 2 == 0

  return h, b
}
