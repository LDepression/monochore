/**
 * @Author: lenovo
 * @Description:
 * @File:  main_generator
 * @Version: 1.0.0
 * @Date: 2023/05/28 16:31
 */

package main

import (
	"fmt"
	"os"
	"path"
)

type MainGenerator struct{}

func (d MainGenerator) Run(opt *Option) error {
	filename := path.Join("./", opt.Output, "main/main.go")
	file, err := os.OpenFile(filename, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0755)
	if err != nil {
		fmt.Printf("open file:%s failed,err:%v", filename, err)
		return err
	}
	defer file.Close()
	/*
		"context"
		proto "etcd/example/grpc_example/hello"
		"google.golang.org/grpc"
		"log"
		"net"

	*/
	fmt.Fprintf(file, "package main\n")
	fmt.Fprintf(file, "import(\n")
	fmt.Fprintf(file, `"net"`)
	fmt.Fprintln(file)
	fmt.Fprintf(file, `"log"`)
	fmt.Fprintln(file)

	fmt.Fprintf(file, `"google.golang.org/grpc"`)
	fmt.Fprintln(file)

	fmt.Fprintf(file, `controllers "etcd/tools/monochore/output/controllers"`)
	fmt.Fprintln(file)
	fmt.Fprintf(file, `proto "etcd/tools/monochore/output/generate"`)
	fmt.Fprintln(file)

	fmt.Fprintf(file, ")\n")
	fmt.Fprintf(file, "var Server = &controllers.Server{}\n")
	fmt.Fprintf(file, "\n\n")

	fmt.Fprintf(file, `var port = "12345"`)
	fmt.Fprintln(file)
	fmt.Fprintf(file, "\n\n")

	fmt.Fprintln(file, `
func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen ：%v", err)
	}
	s := grpc.NewServer()
	proto.RegisterHelloServiceServer(s, Server)
	s.Serve(lis)
}`)
	return nil
}
