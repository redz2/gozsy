package main

import (
	"log"
	"net"
	"net/rpc"
)

type HelloService struct{}

// 方法包含两个可序列化的参数，第二个参数是指针类型
// 返回error类型
func (p *HelloService) Hello(request string, reply *string) error {
	*reply = "hello:" + request
	return nil
}

// 将HelloService注册为rpc服务
func main() {
	rpc.RegisterName("HelloService", new(HelloService))

	listener, err := net.Listen("tcp", ":1234")
	if err != nil {
		log.Fatal("ListenTCP error:", err)
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatal("Accept error:", err)
		}

		go rpc.ServeConn(conn)
	}

}
