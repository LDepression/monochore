package controllers

import (
	"context"
	proto "etcd/tools/monochore/output/generate"
)

// Server 是实现了 HelloServiceServer 接口的结构体
type Server struct {
	proto.UnimplementedHelloServiceServer
}

func (s *Server) SayHello(ctx context.Context, r *proto.HelloRequest) (*proto.HelloResponse, error) {
	// 实现您的方法逻辑
	return nil, nil
}
