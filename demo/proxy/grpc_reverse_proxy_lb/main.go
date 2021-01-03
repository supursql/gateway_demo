package main

import (
	"fmt"
	"gateway_demo/proxy/load_balance"
	"gateway_demo/proxy/proxy"
	"google.golang.org/grpc"
	"log"
	"net"
)

const port = ":50051"

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Printf("failed to listen: %v", err)
	}

	rb := load_balance.LoadBalanceFactory(load_balance.LbWeightRoundRobin)
	rb.Add("127.0.0.1:50055", "40")

	grpcHandler := proxy.NewGrpcLoadBalanceHandler(rb)
	s := grpc.NewServer(grpc.UnknownServiceHandler(grpcHandler))

	fmt.Printf("server listening at %v \n", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

}
