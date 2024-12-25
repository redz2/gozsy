package main

import (
	"flag"
	"fmt"
	"os"
)

var name string
var n int

func init() {
	// 怎样自定义命令源码文件的参数使用说明？必须放在init开头
	flag.CommandLine = flag.NewFlagSet("", flag.ExitOnError)
	flag.CommandLine.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage of %s:\n", "question")
		flag.PrintDefaults()
	}
	// 1. 存储该命令参数值的地址
	// 2. 指定该命令参数的名称
	// 3. 指定在未追加该命令参数时的默认值
	// 4. 该命令参数的简短说明
	flag.StringVar(&name, "name", "everyone", "The greeting object.")
	flag.IntVar(&n, "n", 666, "Number")

}
func main() {
	// 怎样自定义命令源码文件的参数使用说明？对flag.Usage重新赋值，必须放在Parse前面
	// flag.Usage = func() {
	// 	fmt.Fprintf(os.Stderr, "Usage of %s:\n", "question")
	// 	flag.PrintDefaults()
	// }

	flag.Parse()

	fmt.Printf("hello, %s! %d\n", name, n)
}
