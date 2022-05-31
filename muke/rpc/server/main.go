package main

import (
	"log"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
	rpcDemo "studygo/muke/rpc"
)

func main() {
	err := rpc.Register(rpcDemo.DemoService{})
	if err != nil {
		panic(err)
	}

	listener, err := net.Listen("tcp", ":1234")
	if err != nil {
		panic(err)
	}

	for {
		accept, err := listener.Accept()
		if err != nil {
			log.Printf("accept error: %v", err)
			continue
		}

		go jsonrpc.ServeConn(accept)
	}
}
