package main

import (
	"net"
	"strings"
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
	this.conn.Write([]byte(msg + "\r\n"))
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
	} else if len(msg) > 7 && msg[:7] == "rename|" {
		newName := msg[7:]
		this.server.mapLock.Lock()
		//判断是否被占用
		_, ok := this.server.OnlineMap[newName]
		if ok {
			this.sendMsg("当前用户名[" + newName + "]已经被占用")
		} else {
			delete(this.server.OnlineMap, this.Name)
			this.server.OnlineMap[newName] = this
			this.Name = newName
			this.server.mapLock.Unlock()
			this.sendMsg("您已经更新用户名成功")
		}
	} else if len(msg) > 4 && msg[:3] == "to|" {
		pmsg := strings.Split(msg, "|")
		if len(pmsg) != 3 {
			this.sendMsg("消息格式不正确: 正确格式: to|用户名|消息")
		} else {
			toUserName := pmsg[1]
			user, ok := this.server.OnlineMap[toUserName]
			if ok {
				user.sendMsg(this.Name + "对你说:" + pmsg[2])
			} else {
				this.sendMsg("用户不" + toUserName + "存在")
			}
		}

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
