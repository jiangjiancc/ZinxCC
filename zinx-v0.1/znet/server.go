package znet

import (
	"fmt"
	"net"
	"time"

)

type Server struct {

	//服务器的名称
	Name string
	//tcp4 or other
	IPVersion string
	//服务绑定的IP地址
	IP string
	//服务绑定的端口
    Port int
}

//实现ziface.IServer里的全部接口方法

//开启网络服务
func (s *Server) Start()  {
	//1、获取一个TCP的Addr
	addr,err := net.ResolveTCPAddr(s.IPVersion,fmt.Sprintf("%s:%d",s.IP,s.Port))
	if err != nil {
		fmt.Println("resolve tcp addr err:",err)
		return
	}
	//2、监听服务器地址
	listenner,err := net.ListenTCP(s.IPVersion,addr)
	if err != nil {
		fmt.Println("listen",s.IPVersion,"err",err)
		return
	}
	//已经监听成功
	fmt.Println("start Zinx server",s.Name,"succ,now listenning...")
	//3、启动server网络连接业务
	for  {
		//3.1 阻塞等待客户端建立连接请求
		conn,err := listenner.AcceptTCP()
		if err != nil {
			fmt.Println("Accept err",err)
			continue
		}
		//3.2 TODO Server.Stop() 设置服务器最大连接控制，如果超过最大连接，那么则关闭此连接
		//3.3 TODO Server.Start() 处理该连接请求的业务方法，此时应该有handler和conn是绑定的
		//我们这里暂时做一个最大512字节的回显服务
		go func() {
			//不断的循环从客户端获取数据
			for  {
				buf := make([]byte,512)
				cnt,err := conn.Read(buf)
				if err != nil{
					fmt.Println("recv buf err",err)
					continue
				}
				//回显
				if _,err := conn.Write(buf[:cnt]);err != nil{
					fmt.Println("write back buf err",err)
					continue
				}

			}
		}()
	}
}

func (s *Server) Stop()  {
	fmt.Println("[STOP] Zinx server,name",s.Name)
	//TODO  Server.Stop()  将其它需要清理的连接信息或者其它信息，也要一并停止或者清理

	//阻塞，否则主GO退出，listenner的go将会退出
	for  {
		time.Sleep(10 * time.Second)
	}
}

//创建一个服务器句柄
func NewServer(name string)   {

}