package main

import (
	"fmt"
	"log"
	"net/rpc"
)

func main() {
	// 通过 rpc.Dial 拨号 RPC 服务
	client, err := rpc.Dial("tcp", "localhost:1234")
	if err != nil {
		log.Fatal("dialing:", err)
	}

	var reply string
	// client.Call 调用具体的 RPC 方法
	// RPC 服务名字和方法名字

	// 如何更加方便的调用
	err = client.Call("HelloService.Hello", "zy", &reply)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(reply)
}

type HelloServiceClient struct {
	*rpc.Client
}

// var _ HelloServiceInterface = (*HelloServiceClient)(nil)

func DialHelloService(network, address string) (*HelloServiceClient, error) {
	c, err := rpc.Dial(network, address)
	if err != nil {
		return nil, err
	}
	return &HelloServiceClient{Client: c}, nil
}

func (p *HelloServiceClient) Hello(request string, reply *string) error {
	return p.Client.Call("HelloService.Hello", request, reply)
}
