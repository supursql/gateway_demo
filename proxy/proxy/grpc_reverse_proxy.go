package proxy

import (
	"context"
	"gateway_demo/proxy/grpc_proxy"
	"gateway_demo/proxy/load_balance"
	"google.golang.org/grpc"
	"google.golang.org/grpc/encoding"
	"log"
)

func init() {
	encoding.RegisterCodec(grpc_proxy.Codec())
}

func NewGrpcLoadBalanceHandler(lb load_balance.LoadBalance) grpc.StreamHandler {
	return func() grpc.StreamHandler {
		nextAddr, err := lb.Get("")
		if err != nil {
			log.Fatal("get next addr fail")
		}
		director := func(ctx context.Context, fullMethodName string) (context.Context, *grpc.ClientConn, error) {
			c, err := grpc.DialContext(ctx, nextAddr, grpc.WithDefaultCallOptions(grpc.CallContentSubtype("mycodec")), grpc.WithInsecure())
			return ctx, c, err
		}
		return grpc_proxy.TransparentHandler(director)
	}()
}
