/**
 * @Author: lenovo
 * @Description:
 * @File:  main
 * @Version: 1.0.0
 * @Date: 2023/05/27 22:20
 */

package main

import (
	"context"
	proto "etcd/example/grpc_example/hello"
	"google.golang.org/grpc"
	"log"
	"net"
)

const port = ":50051"

type serve struct {
	proto.UnimplementedHelloServiceServer
}

func (s *serve) SayHello(ctx context.Context, in *proto.HelloRequest) (*proto.HelloResponse, error) {
	return &proto.HelloResponse{Reply: "hello" + in.Name}, nil
}
func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen ：%v", err)
	}
	s := grpc.NewServer()
	proto.RegisterHelloServiceServer(s, &serve{})
	s.Serve(lis)
}
