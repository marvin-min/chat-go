package main

import (
	"net"
)

type User struct {
	Name   string
	Addr   string
	C      chan string
	conn   net.Conn
	server *Server
}

func NewUser(conn net.Conn, server *Server) *User {
	addr := conn.RemoteAddr().String()
	user := &User{
		Name:   addr,
		Addr:   addr,
		C:      make(chan string),
		conn:   conn,
		server: server,
	}
	go user.Listener()
	return user
}

// 上线
func (this *User) Online() {
	this.server.mapLock.Lock()
	this.server.OnlineMap[this.Name] = this
	this.server.mapLock.Unlock()
	// 广播上线消息
	this.server.Broadcast(this, "上线了")
}

// 下线
func (this *User) Offline() {
	this.server.mapLock.Lock()
	delete(this.server.OnlineMap, this.Name)
	this.server.mapLock.Unlock()
	this.server.Broadcast(this, "下线")
}

func (this *User) sendMsg(msg string) {
	this.conn.Write([]byte(msg))
}
func (this *User) DoMessage(msg string) {
	if msg == "who" {
		//查询用户
		this.server.mapLock.Lock()
		for _, user := range this.server.OnlineMap {
			olineMsg := "[" + user.Addr + "]" + user.Name + ": 在线"
			this.sendMsg(olineMsg)
		}
		this.server.mapLock.Unlock()
	} else {
		this.server.Broadcast(this, msg)
	}
}

func (this *User) Listener() {
	for {
		msg := <-this.C

		this.conn.Write([]byte(msg + "\n"))
	}
}
