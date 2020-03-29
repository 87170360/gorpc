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

func fillStructField(in interface{}) {
	t := reflect.TypeOf(in)
	v := reflect.ValueOf(in)
	fmt.Println("struct kind", t.Kind())
	if t.Kind() == reflect.Ptr {
		t = t.Elem()
		v = v.Elem()
	}

	for i := 0; i < v.NumField(); i++ {
		typeField := t.Field(i)
		valueField := v.Field(i)
		realValue := valueField.Interface()
		if valueField.CanInterface() && reflect.ValueOf(realValue).IsNil() && reflect.TypeOf(realValue).Kind() == reflect.Ptr {
			//对结构体的field取类型和数值
			fmt.Printf("valueField name:%v type:%v, value:%v\n", typeField.Name, reflect.TypeOf(realValue), reflect.ValueOf(realValue))
			switch reflect.TypeOf(realValue).String() {
			case "*int32":
				var value int32
				valueField.Set(reflect.ValueOf(&value))
			case "*int64":
				var value int64
				valueField.Set(reflect.ValueOf(&value))
			case "*bool":
				var value bool
				valueField.Set(reflect.ValueOf(&value))
			case "*string":
				var value string
				valueField.Set(reflect.ValueOf(&value))
			}
		}
	}
}

//遍历结构体属性
func printStruct(in interface{}) {
	t := reflect.TypeOf(in)
	v := reflect.ValueOf(in)
	//fmt.Println("struct kind", t.Kind())
	if t.Kind() == reflect.Ptr {
		t = t.Elem()
		v = v.Elem()
	}

	for i := 0; i < v.NumField(); i++ {
		typeField := t.Field(i)
		valueField := v.Field(i)
		realValue := valueField.Interface()
		if valueField.CanInterface() {
			/*
				fmt.Printf("Name:%s Type:%s Interface:%v tag:%s \n",
					typeField.Name,
					typeField.Type,
					valueField.Interface(),
					typeField.Tag)
			*/

			//对结构体的field取类型和数值
			fmt.Printf("valueField name:%v type:%v, value:%v, realValueName:%v\n", typeField.Name, reflect.TypeOf(realValue), reflect.ValueOf(realValue), reflect.TypeOf(realValue).String())
		}

		/*
			switch typeField.Type.Kind() {
			case reflect.Ptr:
				fmt.Println("field is ptr")
			default:
				fmt.Println("field not ptr")
			}
		*/
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

	//fmt.Printf("reply:%#v\n", reply)
	printStruct(reply)
	fillStructField(reply)
	println("-------------------------------")
	printStruct(reply)
}
