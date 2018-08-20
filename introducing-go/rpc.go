package main

import(
	"fmt"
	"net"
	"net/rpc"
)

type Server struct { }

func (this *Server) Negate(i int64, reply *int64) error{
	*reply = i * -1
	return nil
}

func server(){
	rpc.Register(new(Server))

	ln, err := net.Listen("tcp", ":9000")
	if err != nil{
		fmt.Println(err)
		return
	}

	for{
		cn, err := ln.Accept()
		if err != nil{
			continue
		}

		go rpc.ServeConn(cn)
	}
}

func client(){
	cn, err := rpc.Dial("tcp", "127.0.0.1:9000")
	if err != nil{
		fmt.Println(err)
		return
	}

	var result int64
	err = cn.Call("Server.Negate", int64(999), &result)

	if err != nil{
		fmt.Println(err)
	}else{
		fmt.Println(result)
	}
}

func main(){
	go server()
	go client()

	var input string
	fmt.Scanln(&input)
}
