package main

import (
	// 模块usetestp的根目录是usetestp
	// package都会从根目录下找，可以在根目录下创建很多package
	"usetestp/pkg"

	// 通过go mod init初始化模块
	// 如何使用本地的模块？
	// 如何使用第三方模块？
	// 如何自己创建一个模块给别人用？

	// 本地模块会去../testp/下找
	"xxx/testp"
)

/* go.mod配置导入本地模块，一定要这种格式: path/module
require xxx/testp v1.1.1
replace xxx/testp => ../testp

*/

func main() {
	// 使用的是package，导入的是package所在的目录
	// 目录名称和package名称一般是一致的

	// 如果不一致，引用时还是需要指定package名称，导入模块过多
	testp.TestP()
	pkg.PrintPackage()
}
