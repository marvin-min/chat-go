package main

import (
	"flag"
	"fmt"
	"net"
)

type Client struct {
	ServerIp   string
	ServerPort int
	Name       string
	conn       net.Conn
}

func NewClient(serverIp string, serverPort int) *Client {
	//创建客户端
	client := &Client{
		ServerIp:   serverIp,
		ServerPort: serverPort,
	}

	//连接server
	addr := fmt.Sprintf("%s:%d", serverIp, serverPort)
	conn, err := net.Dial("tcp", addr)
	if err != nil {
		fmt.Println("Error:", err)
		return nil
	}
	client.conn = conn
	// 返回客户端
	return client
}

var serverIp string
var serverPort int

func init() {
	flag.StringVar(&serverIp, "ip", "127.0.0.1", "服务器ip默认地址127.0.0.1")
	flag.IntVar(&serverPort, "port", 9999, "服务器ip默认端口9999")
}

func main() {
	flag.Parse()
	client := NewClient(serverIp, serverPort)
	if client == nil {
		fmt.Println(">>>>>>连接服务器失败.....")
		return
	}
	fmt.Println("<<<<<<<<<<连接服务器成功.....")
	select {}
}
