package main

import (
	"net"
)

type User struct {
	Name string
	Addr string
	C    chan string
	conn net.Conn
}

func NewUser(conn net.Conn) *User {
	addr := conn.RemoteAddr().String()
	user := &User{
		Name: addr,
		Addr: addr,
		C:    make(chan string),
		conn: conn,
	}
	go user.Listener()
	return user
}

func (this *User) Listener() {
	for {
		msg := <-this.C

		this.conn.Write([]byte(msg + "\n"))
	}
}
