package main

import (
	"log"
	"net"
	"net/rpc"
)

type HelloService struct{}

func (self *HelloService) Hello(request string, reply *string) error {
	*reply = "hello:" + request

	log.Println("got req: ", request)

	return nil
}

func main() {
	rpc.RegisterName("HelloService", new(HelloService))

	listener, err := net.Listen("tcp", ":8888")

	if err != nil {
		log.Fatal("ListenTCP error: ", err)
	}

	rpc.Accept(listener)
}
