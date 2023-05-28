/**
 * @Author: lenovo
 * @Description:
 * @File:  controller_generator
 * @Version: 1.0.0
 * @Date: 2023/05/28 15:23
 */

package main

import (
	"fmt"
	"github.com/emicklei/proto"
	"log"
	"os"
	"path"
)

type CtrlGenerator struct {
	service  *proto.Service
	messages []*proto.Message
	rpc      []*proto.RPC
}

func (d *CtrlGenerator) Run(opt *Option) error {
	reader, err := os.Open(opt.Proto3FileName)
	if err != nil {
		log.Fatalf("open file %s failed err:%v", opt.Proto3FileName, err)
		return err
	}
	defer reader.Close()
	//创建一个解析器
	parser := proto.NewParser(reader)
	definition, err := parser.Parse()
	if err != nil {
		log.Fatalf("parse file %s failed err:%v", opt.Proto3FileName, err)
	}
	proto.Walk(definition,
		proto.WithService(d.handleService),
		proto.WithMessage(d.handleMessage),
		proto.WithRPC(d.handleRPC))
	fmt.Printf("parse protoc secc,rpc:%#v", d.rpc)
	return d.generateRpc(opt)
	return nil
}

func (d *CtrlGenerator) handleService(s *proto.Service) {
	//fmt.Println(s.Name) //HelloService

	d.service = s
}
func (d *CtrlGenerator) handleMessage(m *proto.Message) {

	d.messages = append(d.messages, m)
	//fmt.Println(m.Name) //HelloRequest,HelloResponse

}
func (d *CtrlGenerator) handleRPC(r *proto.RPC) {

	/*
		fmt.Println("rpc begin.................")
		fmt.Println(r.Name)        //SayHello
		fmt.Println(r.RequestType) //HelloRequest
		fmt.Println(r.ReturnsType) //HelloResponse
		fmt.Println("rpc end.................")

	*/
	d.rpc = append(d.rpc, r)
}

func (d *CtrlGenerator) generateRpc(opt *Option) error {
	filename := path.Join("./", opt.Output, "controllers", fmt.Sprintf("%s.go", d.service.Name))
	file, err := os.OpenFile(filename, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0755) //文件不存在就直接创建
	if err != nil {
		fmt.Printf("open file %s: %v", filename, err)
	}
	defer file.Close()
	/*
		生成模板代码
	*/
	fmt.Fprintf(file, "package controllers\n")
	fmt.Fprintf(file, "import(\n")
	fmt.Fprintf(file, `"context"`)
	fmt.Fprintln(file)
	fmt.Fprintf(file, `proto "etcd/tools/monochore/output/generate"`)
	fmt.Fprintln(file)
	fmt.Fprintf(file, ")\n")
	fmt.Fprintf(file, "type Server struct {}")
	fmt.Fprintf(file, "\n\n")

	for _, rpc := range d.rpc {
		fmt.Fprintf(file, "func (s *Server) %s (ctx context.Context,r*hello.%s)(resp*hello.%s,err error){\nreturn\n}\n\n",
			rpc.Name, rpc.RequestType, rpc.ReturnsType)
	}
	return nil
}
