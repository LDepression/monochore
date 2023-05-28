package main

import (
	controllers "etcd/tools/monochore/output/controllers"
	proto "etcd/tools/monochore/output/generate"
	"google.golang.org/grpc"
	"log"
	"net"
)

var Server = &controllers.Server{}

var port = "12345"

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen ：%v", err)
	}
	s := grpc.NewServer()
	proto.RegisterHelloServiceServer(s, Server)
	s.Serve(lis)
}
