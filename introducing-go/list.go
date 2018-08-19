package main

import(
  "fmt"
  "container/list"
  "reflect"
)

func main(){
  newMethodDemo()
  newFunctionDemo()
  listDemo()
}

func newMethodDemo(){
  var l *list.List = new(list.List)
  fmt.Println(l, reflect.TypeOf(l))
}

func newFunctionDemo(){
  var l *list.List = new(list.List)
  fmt.Println(l, reflect.TypeOf(l))
}

func listDemo(){
  var l list.List
  l.PushBack(1)
  l.PushBack(2)
  l.PushBack(3)

  // e is an pointer to list.Element
  // how do data structures implement "template/generic" like functionality
  for e := l.Front(); e != nil;e = e.Next(){
    fmt.Println(e.Value.(int), e, reflect.TypeOf(e))
  }
}
