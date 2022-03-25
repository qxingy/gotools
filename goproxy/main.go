package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"
)

var (
	LocalIp = "0.0.0.0"
	LocalPort = "8080"
)

func init(){
	log.SetPrefix("GOPROXY:")
	log.SetOutput(os.Stdout)
	log.SetFlags(log.Ldate|log.Ltime)
}

func main() {
	log.Println("正在开启...")

	var (
		address = fmt.Sprintf("%s:%s",LocalIp,LocalPort)
	)
	listener , err := net.Listen("tcp" , address)
	if err != nil {
		log.Panicf("端口监听出错!,错误信息:%s" , err)
	}
	defer listener.Close()

	log.Printf("开始监听%s......",address)

	go func (){
		for  {
			log.Printf("正在监听%s,等待请求...",address)
			accept , err := listener.Accept()
			if err != nil {
				log.Printf("连接错误:%s\n",err)
			}

			go Handler(accept)
		}
	}()

	sign := make(chan os.Signal , 0)
	signal.Notify(sign,syscall.SIGINT , syscall.SIGQUIT , syscall.SIGTERM , syscall.SIGKILL)
	select {
		case <- sign :{
			listener.Close()
			log.Println("程序退出")
			os.Exit(1)
		}
	}

}

func Handler(client net.Conn){
	defer func(){
		err := recover()
		log.Println(err)
	}()

	var (
		buffer [2048]byte
	)
	_ , err := client.Read(buffer[:])
	if err != nil {
		log.Panicf("")
	}

}