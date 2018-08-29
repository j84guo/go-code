package main

import(
    "fmt"
    "time"
    "math/rand"
    "strconv"
)

func pinger(c chan string){
    for i:=0; ;i++{
        n := rand.Intn(100)
        c <-"ping" + strconv.Itoa(n)
    }
}

func ponger(c chan string){
    for i:=0; ;i++{
        c <-"pong"
    }
}

func printer(c chan string){
    for{
        msg := <-c
        fmt.Println(msg)
        time.Sleep(time.Second * 1)
    }
}



/*
 * We can restrict a channel to either sending or receiving:
 * func work(c chan <- string) // work may only send to the channel
 * func work(c <- chan string) // work may only receive from the channel
 *
 * A bi-directional channel can be passed to a function which accepts a one-way
 * channel, but the reverse is not true.
 */
func directionDemo(){

}

/*
 * When creating a channel, the make function can be passed a second argument.
 * This integer creates a buffered capacity with a limited capacity. Normally
 * channels are synchronous, each side waits until the other side is ready. A
 * buffered channel is asynchronous, sending or receiving will not wait unless
 * the buffer is full, e.g. c := make(chan string, 1)
 */
func bufferedDemo(){

}

func main(){
    /*
     * A channel type is represented with the chan keyword and the name of the
     * type passed in the channel. The left arrow <- is used to pass data into
     * and receive from the channel.
     *
     * Using a channel synchronizes the two goroutines. When pinger sends a
     * message, it waits until printer is ready to receive, (blocking).
     */
    var c chan string = make(chan string)

    go pinger(c)
    go printer(c)
    go ponger(c)

    var input string
    fmt.Scanln(&input)
}
