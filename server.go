package main

import (
	"fmt"
	"io"
	"net"
	"sync"
	"time"
)

type Server struct {
	Ip   string
	Port int
	//在线用户表
	OnlineMap map[string]*User
	mapLock   sync.RWMutex

	//消息广播
	Message chan string
}

func NewServer(ip string, port int) *Server {
	server := &Server{
		Ip:        ip,
		Port:      port,
		OnlineMap: make(map[string]*User),
		Message:   make(chan string),
	}
	return server
}

// 监听message channel，一旦有消息则广播
func (this *Server) ListenMessage() {
	for {
		msg := <-this.Message
		this.mapLock.Lock()
		for _, cli := range this.OnlineMap {
			cli.C <- msg
		}
		this.mapLock.Unlock()
	}
}

func (this *Server) Broadcast(user *User, msg string) {
	sendMsg := "[" + user.Addr + "]" + " " + user.Name + ":" + msg
	this.Message <- sendMsg
}

// do handler
func (this *Server) Handler(conn net.Conn) {
	user := NewUser(conn, this)
	//将用户加入到上线用户表
	user.Online()
	//监听用户是活跃
	isLive := make(chan bool)
	go func() {
		buf := make([]byte, 4096)
		for {
			n, err := conn.Read(buf)
			if n == 0 {
				user.Offline()
				return
			}
			if err != nil && err != io.EOF {
				fmt.Println("Err:", err)
				return
			}
			msg := string(buf[0 : n-1])

			user.DoMessage(msg)
			isLive <- true
		}
	}()
	//阻塞
	for {
		select {
		case <-isLive:
			//不做任何事情
		case <-time.After(time.Second * 10):
			//已经超时
			//强制关闭当前客户端
			user.sendMsg("你被踢了")
			close(user.C)
			conn.Close()
			// 退出当前Handler
			return
		}
	}
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
	go this.ListenMessage()
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
