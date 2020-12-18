package main

import (
	"fmt"
	"log"
	"net"
)

func main() {
	//1.监听端口
	listener, err := net.Listen("tcp", "0.0.0.0:9090")
	if err != nil {
		log.Fatal("网络不通")
	}
	//2.建立套接字连接
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatal("网络不通")
		}
		go func(conn net.Conn) {
			defer conn.Close()
			for {
				var buf [128]byte
				n, err := conn.Read(buf[:])
				if err != nil {
					fmt.Printf("read from connect failed, err: %v \n", err)
					break
				}
				str := string(buf[:n])
				fmt.Printf("receive from client, data: %v \n", str)
			}
		}(conn)
	}
}
