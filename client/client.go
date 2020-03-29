package main

import (
	"fmt"
	"gorpc"
	"net/rpc"
	"reflect"
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

//遍历结构体属性
func printStruct(in interface{}) {
	t := reflect.TypeOf(in)
	v := reflect.ValueOf(in)
	for i := 0; i < v.NumField(); i++ {
		if v.Field(i).CanInterface() {
			fmt.Printf("Name:%s Type:%s Interface:%v tag:%s \n",
				t.Field(i).Name,
				t.Field(i).Type,
				v.Field(i).Interface(),
				t.Field(i).Tag)
		}
	}
}

func main() {

	args := &common.Args{A: 1, B: 2}
	//reply := &common.NationWar_LoginSyncInfo{}
	//call := NWCall("NWServer.Test", args, reply)

	reply := &common.Reply{}
	call := NWCall("NWServer.Test2", args, reply)

	if call == nil {
		return
	}

	_ = <-call.Done
	if call.Error != nil {
		fmt.Println(call.Error)
		return
	}

	fmt.Printf("reply:%#v\n", reply)
	printStruct(*reply)
}
