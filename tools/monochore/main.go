/**
 * @Author: lenovo
 * @Description:
 * @File:  main
 * @Version: 1.0.0
 * @Date: 2023/05/27 22:59
 */

package main

import (
	"fmt"
	"github.com/urfave/cli/v2"
	"log"
	"os"
)

func main() {
	var opt Option
	app := cli.NewApp()
	app.Flags = []cli.Flag{
		&cli.StringFlag{
			Name:        "f",                 // -f
			Value:       "./hello.proto",     //默认值
			Usage:       "idl filename",      //选项的说明
			Destination: &opt.Proto3FileName, //放在哪里
		},
		&cli.StringFlag{
			Name:        "o",                // -o
			Value:       "./output/",        //默认值
			Usage:       "output directory", //选项的说明 --help
			Destination: &opt.Output,        //放在哪里
		},
		&cli.BoolFlag{
			Name:        "c",                         // -c
			Value:       false,                       //默认值
			Usage:       "genetate grpc client code", //选项的说明
			Destination: &opt.GenClientCode,          //放在哪里
		},
		&cli.BoolFlag{
			Name:        "s",                        // -s
			Value:       false,                      //默认值
			Usage:       "genetate grpc serve code", //选项的说明
			Destination: &opt.GenServerCode,         //放在哪里
		},
	}
	app.Action = func(c *cli.Context) error {
		//命令行程序的代码在这里实现

		err := genMgr.Run(&opt)
		return err
	}

	err := app.Run(os.Args) //这里其实调用的就是Action
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("code generation successfully")
}
