package main

import (
	"flag"
	"fmt"
)

func main() {
	// 获取命令行参数
	// args := os.Args

	// 定义一个命令行参数
	name := flag.String("name", "zhouyi", "give a name") //返回*string
	// 解析命令行参数
	flag.Parse()
	fmt.Println(*name)
}
