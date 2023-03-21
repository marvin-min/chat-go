package main

import (
	"flag"
	"fmt"
	"io"
	"net"
	"os"
)

type Client struct {
	ServerIp   string
	ServerPort int
	Name       string
	conn       net.Conn
	flag       int
}

func NewClient(serverIp string, serverPort int) *Client {
	//创建客户端
	client := &Client{
		ServerIp:   serverIp,
		ServerPort: serverPort,
		flag:       999,
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

func (this *Client) menu() bool {
	var flag int
	fmt.Println("1. 群聊\r\n2.私聊\r\n3.更新用户名\r\n0.退出")
	fmt.Scanln(&flag)
	if flag >= 0 && flag <= 3 {
		this.flag = flag
		return true
	} else {
		fmt.Println(">>>>请输入0~3之间有效的菜单")
		return false
	}
}
func (client *Client) UpdateName() bool {
	fmt.Println("》》》》》请输入用户名:")
	fmt.Scanln(&client.Name)
	msg := "rename|" + client.Name + "\n"
	_, err := client.conn.Write([]byte(msg))
	if err != nil {
		fmt.Println("conn.Write err:", err)
		return false
	}
	return true
}

func (client *Client) Run() {
	for client.flag != 0 {
		for client.menu() != true {

		}
		switch client.flag {
		case 1:
			//群聊
			fmt.Println("群聊-----")
			break
		case 2:
			//私聊
			fmt.Println("私聊-----")
			break
		case 3:
			//改名
			client.UpdateName()
			break
		}
	}
}

var serverIp string
var serverPort int

func init() {
	flag.StringVar(&serverIp, "ip", "127.0.0.1", "服务器ip默认地址127.0.0.1")
	flag.IntVar(&serverPort, "port", 9999, "服务器ip默认端口9999")
}

// 处理server回应的消息，直接显示到终端
func (client *Client) DealResponse() {
	io.Copy(os.Stdout, client.conn)
}

func main() {
	flag.Parse()
	client := NewClient(serverIp, serverPort)
	if client == nil {
		fmt.Println(">>>>>>连接服务器失败.....")
		return
	}
	go client.DealResponse()
	fmt.Println("<<<<<<<<<<连接服务器成功.....")
	client.Run()
}
