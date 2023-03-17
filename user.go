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

func (this *User) DoMessage(msg string) {
	this.server.Broadcast(this, msg)
}

func (this *User) Listener() {
	for {
		msg := <-this.C

		this.conn.Write([]byte(msg + "\n"))
	}
}
