package main

import (
	"net"
	"strings"
)

type User struct {
	Name string
	Addr string
	C    chan string
	conn net.Conn

	server *Server
}

//创建一个用户的api
func NewUser(conn net.Conn, server *Server) *User {
	userAddr := conn.RemoteAddr().String()
	user := &User{
		Name:   userAddr,
		Addr:   userAddr,
		C:      make(chan string),
		conn:   conn,
		server: server,
	}
	go user.ListenMessage()
	return user

}

//用户上线的业务
func (this *User) Online() {

	//用户上线 将用户加入onlinemap中
	this.server.mapLock.Lock()
	this.server.OnLineMap[this.Name] = this
	this.server.mapLock.Unlock()
	//广播用户上线的消息
	this.server.BroadCast(this, "已上线")
}

//用户下线的业务
func (this *User) Offline() {
	//用户上线 将用户加入onlinemap中
	this.server.mapLock.Lock()
	delete(this.server.OnLineMap, this.Name)
	this.server.mapLock.Unlock()
	//广播用户上线的消息
	this.server.BroadCast(this, "xia线")
}

//给当前user对应的客户端发送消息
func (this *User) SendMsg(msg string) {
	this.conn.Write([]byte(msg))
}

//处理用户消息
func (this *User) DoMessage(msg string) {

	if msg == "who" {
		//查询当前在线用户有哪些
		this.server.mapLock.Lock()
		for _, user := range this.server.OnLineMap {
			onlineMsg := "[" + user.Name + "]" + ":" + "online。。。\n"
			this.SendMsg(onlineMsg)
		}
		this.server.mapLock.Unlock()
	} else if len(msg) > 7 && msg[:7] == "rename|" {
		newName := strings.Split(msg, "|")[1]
		_, ok := this.server.OnLineMap[newName]
		if ok {
			this.SendMsg("当前用户名被使用")
		} else {
			this.server.mapLock.Lock()
			delete(this.server.OnLineMap, this.Name)
			this.server.OnLineMap[newName] = this
			this.server.mapLock.Unlock()

			this.Name = newName
			this.SendMsg("您已更新用户名：" + this.Name + "\n")
		}
	} else if len(msg) > 4 && msg[:3] == "to|" {

		tempStr := strings.Split(msg, "|")

		if tempStr[1] == "" {
			this.SendMsg("消息格式不正确，请使用 to|张三|你好啊  格式。\n")
			return
		}
		remoteUser, ok := this.server.OnLineMap[tempStr[1]]
		if !ok {
			this.SendMsg("该用户名不存在！")
			return
		}
		if tempStr[2] == "" {
			this.SendMsg("消息内容为空，请重发")
			return
		}
		remoteUser.SendMsg(this.Name + "对您说：" + tempStr[2])
	} else {

		this.server.BroadCast(this, msg)
	}
}

//监听当前user channel的方法 一旦有消息 就直接发送给对应的客户端
func (this *User) ListenMessage() {
	for {
		msg := <-this.C
		this.conn.Write([]byte(msg + "\n"))
	}
}
