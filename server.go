package main

import (
	"fmt"
	"net"
)

type Server struct {
	Ip   string
	Port int
}

func NewServer(ip string, port int) *Server {
	server := &Server{
		Ip:   ip,
		Port: port,
	}
	return server
}

// do handler
func (this *Server) Handler(conn net.Conn) {
	fmt.Println("连接建立成功")
}
func (this *Server) Start() {
	fmt.Println("启动server")
	//socket listen
	list, err := net.Listen("tcp", fmt.Sprintf("%s:%d", this.Ip, this.Port))
	if err != nil {
		fmt.Println(err)
	}
	//close
	defer list.Close()
	fmt.Println("准备开始接受请求")
	//socket listen
	for {
		//accept
		conn, err := list.Accept()
		if err != nil {
			fmt.Println("Listener accept err:", err)
			continue
		}
		go this.Handler(conn)
	}
}
