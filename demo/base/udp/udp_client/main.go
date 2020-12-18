package main

import (
	"fmt"
	"net"
)

func main() {

	/**
	func DialUDP(net string, laddr, raddr *UDPAddr) (*UDPConn, error)
	DialTCP在网络协议net上连接本地地址laddr和远端地址raddr。net必须是"udp"、"udp4"、"udp6"；如果laddr不是nil，将使用它作为本地地址，否则自动选择一个本地地址。
	*/
	conn, err := net.DialUDP("udp", nil, &net.UDPAddr{
		IP:   net.IPv4(127, 0, 0, 1),
		Port: 9090,
	})

	if err != nil {
		fmt.Printf("connect failed, err : %v \n", err)
		return
	}

	for i := 0; i < 100; i++ {
		_, err = conn.Write([]byte("hello server!"))
		if err != nil {
			fmt.Printf("send data failed, err : %v\n", err)
			return
		}

		result := make([]byte, 1024)
		n, remoteAddr, err := conn.ReadFromUDP(result)
		if err != nil {
			fmt.Printf("recevie data failed, err : %v\n", err)
			return
		}
		fmt.Printf("recevie from addr : %v data : %v \n", remoteAddr, string(result[:n]))
	}
}
