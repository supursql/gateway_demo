package grpc_proxy

import (
	"github.com/golang/protobuf/proto"
	"google.golang.org/grpc/encoding"
)

// Codec returns a proxying grpc.Codec with the default protobuf codec as parent.
//
// See CodecWithParent.
// 构建输出函数
func Codec() encoding.Codec {
	return &myCodec{}
}

type myCodec struct {
}

type frame struct {
	payload []byte
}

//构建原始字节解码器
func (c *myCodec) Marshal(v interface{}) ([]byte, error) {
	out, ok := v.(*frame)
	if !ok {
		return proto.Marshal(v.(proto.Message))
	}
	return out.payload, nil

}

func (c *myCodec) Unmarshal(data []byte, v interface{}) error {
	dst, ok := v.(*frame)
	if !ok {
		return proto.Unmarshal(data, v.(proto.Message))
	}
	dst.payload = data
	return nil
}

func (c *myCodec) Name() string {
	return "mycodec"
}
