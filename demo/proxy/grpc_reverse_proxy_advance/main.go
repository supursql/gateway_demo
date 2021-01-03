package main

import (
	"fmt"
	"gateway_demo/proxy/grpc_interceptor"
	"gateway_demo/proxy/grpc_proxy"
	"gateway_demo/proxy/load_balance"
	proxy2 "gateway_demo/proxy/proxy"
	"gateway_demo/proxy/public"
	"google.golang.org/grpc"
	"google.golang.org/grpc/encoding"
	"log"
	"net"
	"time"
)

const port = ":50051"

func init() {
	encoding.RegisterCodec(grpc_proxy.Codec())
	fmt.Println(encoding.GetCodec("mycodec").Name())
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	rb := load_balance.LoadBalanceFactory(load_balance.LbWeightRoundRobin)
	rb.Add("127.0.0.1:50055", "40")

	counter, _ := public.NewFlowCountService("local_app", time.Second)
	grpcHandler := proxy2.NewGrpcLoadBalanceHandler(rb)
	s := grpc.NewServer(
		grpc.ChainStreamInterceptor(
			grpc_interceptor.GrpcAuthStreamInterceptor,
			grpc_interceptor.GrpcFlowCountStreamInterceptor(counter)),
		grpc.UnknownServiceHandler(grpcHandler))

	fmt.Printf("server listening at %v\n", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
