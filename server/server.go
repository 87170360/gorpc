package main

import (
	"fmt"
	proto "github.com/golang/protobuf/proto"
	"gorpc"
	"net"
	"net/rpc"
	"time"
)

type NWServer int

func (this *NWServer) Test(args *common.Args, reply *common.NationWar_LoginSyncInfo) (err error) {
	reply.ZoneID = proto.Int32(0)
	var value int32
	value = 0
	reply.ZoneNum = proto.Int32(value)
	reply.RunState = proto.Int32(1)

	fmt.Printf("rpc server func Test, args:%#v, reply%#v\n", args, reply)
	return
}

func (this *NWServer) Test2(args *common.Args, reply *common.Reply) (err error) {
	reply.A = proto.Int32(0)
	var value int32
	value = 1
	reply.B = proto.Int32(value)

	value2 := 0
	reply.C = &value2

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
