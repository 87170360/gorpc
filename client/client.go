package main

import (
	"fmt"
	"gorpc"
	"net/rpc"
)

func getRPCClient() (client *rpc.Client, err error) {
	client, err = rpc.Dial("tcp", common.ServerAddr)
	if err != nil {
		fmt.Println("dialing:", err)
	}
	return
}

func NWCall(funcName string, args interface{}, reply interface{}) *rpc.Call {
	client, err := getRPCClient()
	if err != nil {
		fmt.Println("call:", err)
		return nil
	}

	call := client.Go(funcName, args, reply, nil)
	return call
	//replyCall := <-call.Done
}

func main() {

	args := &common.Args{A: 1, B: 2}
	reply := &common.Args{}
	call := NWCall("NWServer.Test", args, reply)

	if call == nil {
		return
	}

	_ = <-call.Done
	if call.Error != nil {
		fmt.Println(call.Error)
		return
	}

	fmt.Printf("args:%#v, reply:%#v\n", args, reply)
}
