/**
 * @Author: lenovo
 * @Description:
 * @File:  dir_generator
 * @Version: 1.0.0
 * @Date: 2023/05/28 13:06
 */

package main

import (
	"fmt"
	"os"
	"path"
)

var AllDirList = []string{
	"controllers",
	"idl",
	"main",
	"scripts",
	"conf",
	"app/router",
	"app/config",
	"model",
	"generate",
}

type DirGenerator struct {
	dirList []string
}

func (d *DirGenerator) Run(opt *Option) error {

	for _, dir := range AllDirList {
		fullDir := path.Join(opt.Output, dir)
		if err := os.MkdirAll(fullDir, 0755); err != nil {
			fmt.Printf("mkdir dir %s failed,error: %v", dir, err)
			return err
		}
	}
	return nil
}

//func init() {
//	dir := &DirGenerator{
//		dirList: AllDirList,
//	}
//	Register("dir generator", dir)
//}
