package main

import (
	"context"
	"flag"
	"fmt"
	"gateway_demo/demo/proxy/grpc_server_client/proto"
	"github.com/golang/glog"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"google.golang.org/grpc"
	"net/http"
)

var (
	serverAddr         = ":8001"
	grpcServerEndpoint = flag.String("grpc-server-endpoint", "localhost:50055", "gRPC server endpoint")
)

func run() error {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	//可以理解为每个rs都需要持续跟下游建立连接
	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithInsecure()}
	err := proto.RegisterEchoHandlerFromEndpoint(ctx, mux, *grpcServerEndpoint, opts)
	if err != nil {
		return err
	}
	return http.ListenAndServe(serverAddr, mux)
}

func main() {
	flag.Parse()
	defer glog.Flush()

	fmt.Println("server listening at ", serverAddr)
	if err := run(); err != nil {
		glog.Fatal(err)
	}
}
