package main

import (
	"log"
	"net"
	"net/rpc/jsonrpc"
	rpcDemo "studygo/muke/rpc"
)

func main() {
	conn, err := net.Dial("tcp", ":1234")
	if err != nil {
		panic(err)
	}

	client := jsonrpc.NewClient(conn)
	var result float64
	err = client.Call("DemoService.Div", rpcDemo.Args{A: 10, B: 3}, &result)
	if err != nil {
		log.Printf("err is %v", err)
	} else {
		log.Printf("result %v ", result)
	}

}
