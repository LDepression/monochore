/**
 * @Author: lenovo
 * @Description:
 * @File:  generator_mgr
 * @Version: 1.0.0
 * @Date: 2023/05/27 23:13
 */

package main

import "errors"

// 这里的genMgr不可以使用init创建，因为在其他.go文件里也使用了init函数，不一定哪个init先执行
// 所以这里使用全局变量创建，全局变量最先执行
var genMgr = &GeneratorMgr{
	GeneratorMap: make(map[string]Generator),
}

type GeneratorMgr struct {
	GeneratorMap map[string]Generator
}

func (g *GeneratorMgr) Run(opt *Option) error {
	//遍历所有的map，然后执行对应的Generator就好了
	for _, gen := range g.GeneratorMap {
		err := gen.Run(opt)
		return err
	}
	return nil
}

func Register(name string, gen Generator) error {
	_, ok := genMgr.GeneratorMap[name]
	if ok {
		return errors.New("generator already registered")
	}
	genMgr.GeneratorMap[name] = gen
	return nil
}
