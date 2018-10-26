package main

import "fmt"

func first() {
	fmt.Println("first")
}

func second() {
	fmt.Println("second")
}

/*
go provides a special keyword defer which schedules a function call at the end
of the calling function, it is often used when resources need to be freed

(go does not do RAII using constructorrs/destructors like C++)
- defer functions are run before each return statement (if there are multiple)
- and even if a panic occurs
- supposedly defer can be used to keep the closing logic of a function close to
  the opening logic
*/
func main() {
	deferDemo()
	recoverDemo()
}

func deferDemo() {
	defer second()
	first()
}

/*
calling the panic() function causes a runtime error
these can be handled by the recover() function

recover() stops the panic and returns the value that was passed to panic()

since panic() immediately stops execution of a function, recover() must be
deferred
*/
func recoverDemo() {
	defer func() {
		str := recover()
		fmt.Println(str)
	}()

	panic("panic recovered")
}
