package main

import (
	"fmt"
	"gateway_demo/demo/base/unpack"
	"net"
)

func main() {

	/**
	func Dial(network, address string) (Conn, error)
	在网络network上连接地址address，并返回一个Conn接口。可用的网络类型有：
	"tcp"、"tcp4"、"tcp6"、"udp"、"udp4"、"udp6"、"ip"、"ip4"、"ip6"、"unix"、"unixgram"、"unixpacket"
	对TCP和UDP网络，地址格式是host:port或[host]:port，参见函数JoinHostPort和SplitHostPort。
	*/
	conn, err := net.Dial("tcp", "localhost:9090")
	defer conn.Close()
	if err != nil {
		fmt.Printf("connect failed, err : %v \n", err.Error())
		return
	}
	unpack.Encode(conn, "hello world 0!!!")
}
