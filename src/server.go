package main

import (
	"fmt"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
	"server"
	"time"
	"strings"
)


type Arith int

func (t *Arith) Receive(args *string, reply *int64) error {
	fmt.Println("用户名："+*args)
	fmt.Println("开始进行客户端验证...")
	if strings.Compare(*args,server.USERNAME)==0{
		*reply = time.Now().UnixNano()
	}else {
		fmt.Println("用户未能验证通过")
		*reply = 0
	}
	return nil
}

func main() {

	server.GetServerAddr()
	arith := new(Arith)
	rpc.Register(arith)

	tcpAddr, err := net.ResolveTCPAddr("tcp", server.SERVERIP+":"+server.SERVERPORT)
	if err != nil {
		fmt.Println(err)
	}

	listener, err := net.ListenTCP("tcp", tcpAddr)

	fmt.Println("正在监听端口",server.SERVERPORT)
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("未能监听到请求")
			return
		}
		go jsonrpc.ServeConn(conn)
	}
}