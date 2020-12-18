package main

import (
	"fmt"
	"net"
)

func main() {
	listener, err := net.ListenUDP("udp", &net.UDPAddr{
		IP:   net.IPv4(0, 0, 0, 0),
		Port: 9090,
	})

	if err != nil {
		fmt.Printf("listen failed, err : %v \n", err)
		return
	}

	for {
		var data [1024]byte
		/**
		func (c *UDPConn) ReadFromUDP(b []byte) (n int, addr *UDPAddr, err error)
		ReadFromUDP从c读取一个UDP数据包，将有效负载拷贝到b，返回拷贝字节数和数据包来源地址。
		ReadFromUDP方法会在超过一个固定的时间点之后超时，并返回一个错误。
		*/
		n, addr, err := listener.ReadFromUDP(data[:])
		if err != nil {
			fmt.Printf("read failed from addr : %v,  err : %v \n", addr, err)
			break
		}

		go func() {
			fmt.Printf("addr : %v data : %v count : %v \n", addr, string(data[:n]), n)
			/**
			func (c *UDPConn) WriteToUDP(b []byte, addr *UDPAddr) (int, error)
			WriteToUDP通过c向地址addr发送一个数据包，b为包的有效负载，返回写入的字节。
			WriteToUDP方法会在超过一个固定的时间点之后超时，并返回一个错误。在面向数据包的连接上，写入超时是十分罕见的。
			*/
			_, err = listener.WriteToUDP([]byte("recevied success!"), addr)
			if err != nil {
				fmt.Printf("write failed, err : %v \n", err)
			}
		}()

	}
}
