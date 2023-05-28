/**
 * @Author: lenovo
 * @Description:
 * @File:  interface
 * @Version: 1.0.0
 * @Date: 2023/05/27 22:48
 */

package main

type Generator interface {
	Run(opt *Option) error
}
