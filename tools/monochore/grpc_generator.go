/**
 * @Author: lenovo
 * @Description:
 * @File:  grpc_generator
 * @Version: 1.0.0
 * @Date: 2023/05/28 13:30
 */

package main

import (
	"fmt"
	"os/exec"
)

type GrpcGenerator struct{}

func (d *GrpcGenerator) Run(opt *Option) error {
	////protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative *.proto;
	//cmd := exec.Command("protoc", "--go_out=. --go_opt=paths=source_relative",
	//	"--go-grpc_out=. --go-grpc_opt=paths=source_relative", opt.Proto3FileName)
	outputParams := fmt.Sprintf("%s/generate", opt.Output)
	cmd := exec.Command("protoc", "--go_out", outputParams, "--go_opt", "paths=source_relative", "--go-grpc_out", outputParams, "--go-grpc_opt", "paths=source_relative", opt.Proto3FileName)
	output, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println("grpc generated error:", err)
		fmt.Println("grpc generated output:", string(output))
		return err
	}
	return nil
}

//func init() {
//	dir := &GrpcGenerator{}
//	Register("grpc generator", dir)
//}
