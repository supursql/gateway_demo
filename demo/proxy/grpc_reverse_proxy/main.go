package main

import (
	"context"
	"fmt"
	"gateway_demo/proxy/grpc_proxy"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/encoding"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	"log"
	"net"
	"strings"
)

const (
	port = ":50051"
)

func init() {
	encoding.RegisterCodec(grpc_proxy.Codec())
	fmt.Println(encoding.GetCodec("mycodec").Name())
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	director := func(ctx context.Context, fullMethodName string) (context.Context, *grpc.ClientConn, error) {
		// 拒绝某些特殊请求
		if strings.HasPrefix(fullMethodName, "/com.example.internal.") {
			return ctx, nil, status.Errorf(codes.Unimplemented, "Unknown method")
		}
		c, err := grpc.DialContext(ctx, "localhost:50055", grpc.WithDefaultCallOptions(grpc.CallContentSubtype("mycodec")), grpc.WithInsecure())
		md, _ := metadata.FromIncomingContext(ctx)
		outCtx, _ := context.WithCancel(ctx)
		outCtx = metadata.NewOutgoingContext(outCtx, md)
		return ctx, c, err
	}

	s := grpc.NewServer(
		grpc.UnknownServiceHandler(grpc_proxy.TransparentHandler(director)))
	fmt.Printf("server listening at %v\n", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
