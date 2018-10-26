package main

import (
	"encoding/gob"
	"fmt"
	"net"
)

/*
* the net.Listen function creates a listener
*
* type Listener interface{
* 	  Accept()(c Conn, err error) // next connection
*     Close() error // close socket
*	  Addr() Addr // network address
* }
*
* we use net.Accept to obtain new connections
* net.Conn implements io.Reader and io.Writer interfaces
 */
func server() {
	ln, err := net.Listen("tcp", "127.0.0.1:8000")
	if err != nil {
		fmt.Println(err)
		return
	}

	for {
		cn, err := ln.Accept()
		if err != nil {
			fmt.Println(err)
			continue
		}

		// goroutine
		go handleConnection(cn)
	}
}

func handleConnection(cn net.Conn) {
	var msg string
	err := gob.NewDecoder(cn).Decode(&msg)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(msg)
	}

	cn.Close()
}

func client(msg string) {
	cn, err := net.Dial("tcp", "127.0.0.1:8000")
	if err != nil {
		fmt.Println(err)
		return
	}

	err = gob.NewEncoder(cn).Encode(msg)
	if err != nil {
		fmt.Println(err)
	}

	cn.Close()
}

func main() {
	go server()

	for {
		var msg string
		fmt.Scanln(&msg)
		client(msg)
	}
}
