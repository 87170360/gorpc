package main

import (
	"fmt"
	"gorpc"
	"net"
	"net/rpc"
	"time"
)

type NWServer int

func (this *NWServer) Test(args *common.Args, reply *common.Reply) (err error) {
	var value1, value2 int
	value1 = 1
	reply.A = &value1

	value2 = 0
	reply.B = &value2

	fmt.Printf("rpc server func Test, args:%#v, reply%#v\n", args, reply)
	return
}

func init() {
	rpc.Register(new(NWServer))
}

func Start() {
	listener, err := net.Listen("tcp", common.ServerAddr)
	if err != nil {
		fmt.Println("RPC server listen error", err)
		return
	}
	fmt.Println("RPC server listening on:", common.ServerAddr)

	go rpc.Accept(listener)
}

func main() {
	Start()
	for {
		time.Sleep(1 * time.Second)
	}
}
