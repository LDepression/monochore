/**
 * @Author: lenovo
 * @Description:
 * @File:  enter
 * @Version: 1.0.0
 * @Date: 2023/05/28 14:42
 */

package main

func init() {
	grpcDir := &GrpcGenerator{}
	//dir := &DirGenerator{
	//	dirList: AllDirList,
	//}
	//Register("dir generator", dir)
	Register("grpc generator", grpcDir)
}
