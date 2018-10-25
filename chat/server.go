package main

import (
    "fmt"
    "net"
    "os"
)

const (
    HOST = "localhost"
    PORT = "3333"
    TYPE = "tcp"
)

func main() {
    fmt.Println("A concurrent TCP server using Goroutines")

    ls, err := net.Listen(TYPE, HOST + ":" + PORT)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error listening: %s", err.Error())
        os.Exit(1)
    }

    defer ls.Close()
    fmt.Println("Listening locally on port " + PORT)

    for {
        cn, err := ls.Accept()
        if err != nil {
            fmt.Fprintf(os.Stderr, "Error accepting: %s", err.Error())
            os.Exit(1)
        }

        // new goroutine per request, how does this compare to select()
        go handle(cn)
    }
}

func handle(cn net.Conn) {
    defer cn.Close()

    buf := make([]byte, 1024)
    _, err := cn.Read(buf)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error reading: %s", err.Error())
    }

    cn.Write([]byte("Hello world!\n"))
}

