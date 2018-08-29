package main

import(
    "fmt"
    "time"
    "reflect"
)

func main(){
    selectDemo()

    c1 := make(chan string)
    c2 := make(chan string)
    timeoutDemo(c1, c2)
}

func timeoutDemo(c1 chan string, c2 chan string){
    select{
    case msg1 := <-c1:
        fmt.Println("msg1", msg1)
    case msg2 := <-c2:
        fmt.Println("msg2", msg2)

    /*
     * time.After creates a channel and sends the current time to it after one
     * second
     */
    case <- time.After(time.Second):
        fmt.Println("timeout")
    }
}

/*
 * Go has a special select statement that works like a switch but for channels,
 * picking a channel to read or write to.
 *
 * Select picks the first channel which is ready. If no channels are ready it
 * blocks. If multiple are ready it randomly picks one.
 *
 * The default block is executed if neither channel is immediately ready.
 */
func selectDemo(){
    c1 := make(chan string)
    c2 := make(chan string)

    go func(){
        for{
            c1 <- "from 1"
            time.Sleep(time.Second * 2)
        }
    }()

    go func(){
        for{
            c2 <- "from 2"
            time.Sleep(time.Second * 3)
        }
    }()

    go func(){
        for{
            /*
             * print the first channel which has data
             */
            select{
            case msg1 := <-c1:
                fmt.Println(msg1, reflect.TypeOf(msg1))
            case msg2 := <-c2:
                fmt.Println(msg2, reflect.TypeOf(msg2))
            }
        }
    }()

    var input string
    fmt.Scanln(&input)
}
