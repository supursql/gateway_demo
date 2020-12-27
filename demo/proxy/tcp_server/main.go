package main

import (
	"context"
	"fmt"
	"gateway_demo/proxy/tcp_proxy"
	"log"
	"net"
)

var (
	addr = ":7002"
)

type tcpHandler struct {
}

func (t *tcpHandler) ServeTCP(ctx context.Context, src net.Conn) {
	src.Write([]byte("tcpHandler\n"))
}

func main() {
	//tcp 服务测试
	log.Println("Starting tcpserver at " + addr)
	tcpServ := tcp_proxy.TcpServer{
		Addr:    addr,
		Handler: &tcpHandler{},
	}
	fmt.Println("Starting tcp_server at " + addr)
	tcpServ.ListenAndServe()

}
