/**
 * @Author: lenovo
 * @Description:
 * @File:  main
 * @Version: 1.0.0
 * @Date: 2023/05/27 22:32
 */

package main

import (
	"context"
	proto "etcd/example/grpc_example/hello"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const (
	address     = "localhost:50051"
	defaultName = "world"
)

func main() {
	conn, err := grpc.Dial("127.0.0.1:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return
	}
	defer conn.Close()

	c := proto.NewHelloServiceClient(conn)
	r, err := c.SayHello(context.Background(), &proto.HelloRequest{Name: defaultName})
	if err != nil {
		return
	}
	fmt.Println(r.Reply)
}
