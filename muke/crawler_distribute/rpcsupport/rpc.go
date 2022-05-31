package rpcsupport

import (
	"log"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
)

func ServeRpc(host string, service interface{}) error {
	err := rpc.Register(service)
	if err != nil {
		return nil
	}

	listener, err := net.Listen("tcp", host)
	if err != nil {
		panic(err)
	}

	log.Printf("Listening on %s ", host)
	for {
		accept, err := listener.Accept()
		if err != nil {
			log.Printf("accept error: %v", err)
			continue
		}

		go jsonrpc.ServeConn(accept)
	}
	return nil
}

func NewClient(host string) (*rpc.Client, error) {
	conn, err := net.Dial("tcp", host)
	if err != nil {
		return nil, err
	}

	return jsonrpc.NewClient(conn), nil
}
