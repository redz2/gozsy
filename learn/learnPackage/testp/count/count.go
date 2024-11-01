package count

import "fmt"

// 小写开头，私有变量，只能包内使用
var pc int

// 导入package时，会自动调用init
func init() {
	fmt.Println("func init is running")
	pc = 2
}

// 大写开头，包外可用
func PrintPc() {
	fmt.Println(pc)
}
