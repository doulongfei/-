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
	//	在线用户的列表
	OnLineMap map[string]*User
	mapLock   sync.RWMutex
	//消息广播的channel
	Message chan string
}

//创建一个server接口
func NewServer(ip string, port int) *Server {
	server := &Server{
		Ip:        ip,
		Port:      port,
		OnLineMap: make(map[string]*User),
		Message:   make(chan string),
	}
	return server
}

//启动服务器的接口
func (this *Server) Start() {
	//socket listen
	listener, err := net.Listen("tcp", fmt.Sprintf("%s:%d", this.Ip, this.Port))
	if err != nil {
		fmt.Println("net,listen err", err)
		return
	}
	defer listener.Close()

	//启动监听message的goroutine
	go this.ListeMessage()
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("listern accept err", err)
			continue
		}
		go this.Handler(conn)
	}
}

func (this *Server) Handler(conn net.Conn) {
	//...当前连接的业务
	fmt.Println("连接成功。。。")
	user := NewUser(conn, this)
	////用户上线 将用户加入onlinemap中
	//this.mapLock.Lock()
	//this.OnLineMap[user.Name] = user
	//this.mapLock.Unlock()
	////广播用户上线的消息
	//this.BroadCast(user, "已上线")
	user.Online()

	isLive := make(chan bool)
	go func() {
		buf := make([]byte, 4096)
		for {
			n, err := conn.Read(buf)
			if n == 0 {
				user.Offline()
				//this.BroadCast(user,"xiaXian")
				return
			}
			if err != nil && err != io.EOF {
				fmt.Println("conn read err", err)
				return
			}
			msg := string(buf[:n-1])
			//this.BroadCast(user,msg)
			//	用户针对msg进行消息处理
			user.DoMessage(msg)

			isLive <- true
		}

	}()
	//	当前handler阻塞
	for {

		select {
		case <-isLive:

		case <-time.After(time.Second * 10):

			user.SendMsg("你被踢出聊天")

			close(user.C)

			conn.Close()

			return
		}
	}

}

//广播消息
func (this *Server) BroadCast(user *User, msg string) {
	sendMsg := "{" + user.Addr + "}" + user.Name + ":" + msg
	this.Message <- sendMsg

}

//监听message广播消息channel的goroutine 一旦有消息就发送给全部在线的user
func (this *Server) ListeMessage() {
	for {
		msg := <-this.Message
		//将msg发送给全部的在线user
		this.mapLock.Lock()
		for _, cli := range this.OnLineMap {
			cli.C <- msg
		}
		this.mapLock.Unlock()
	}
}
