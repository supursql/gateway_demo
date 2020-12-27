package main

import (
	"context"
	"fmt"
	"gateway_demo/proxy/load_balance"
	"gateway_demo/proxy/proxy"
	"gateway_demo/proxy/tcp_middleware"
	"gateway_demo/proxy/tcp_proxy"
	"net"
)

var (
	addr = ":2002"
)

type tcpHandler struct {
}

func (t *tcpHandler) ServeTCP(ctx context.Context, src net.Conn) {
	src.Write([]byte("tcpHandler\n"))
}

func main() {
	//tcp 服务测试
	//log.Println("Starting tcpserver at " + addr)
	//tcpServ := tcp_proxy.TcpServer{
	//	Addr: addr,
	//	Handler: &tcpHandler{},
	//}
	//fmt.Println("Starting tcp_server at " + addr)
	//tcpServ.ListenAndServe()

	//代理测试
	rb := load_balance.LoadBalanceFactory(load_balance.LbWeightRoundRobin)
	rb.Add("127.0.0.1:7002", "100")
	proxy := proxy.NewTcpBalanceReverseProxy(&tcp_middleware.TcpSliceRouterContext{}, rb)
	tcpServ := tcp_proxy.TcpServer{Addr: addr, Handler: proxy}
	fmt.Println("Starting tcp_proxy at " + addr)
	tcpServ.ListenAndServe()
}
