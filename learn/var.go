package learn

import (
	"flag"
	"fmt"
)

// 声明变量的几种方式
// 1，

func LearnVar() {
	var name string
	flag.StringVar(&name, "name", "everyone", "The greeting object.")

	// go语言中的类型推断，短变量声明的用法（只能在函数体内部使用短变量声明）
	// 类型推断：编程语言在编译期自动解释表达式类型的能力
	// var name string
	// var name = "Robert"
	// name := "Robert"

	flag.Parse()
	fmt.Printf("Hello, %v!\n", name)

	// 1，go语言的类型推断能带来哪些好处？
	// 代码重构：不改变某个程序和外界的任何交互方式和规则，只改变内部实现

	// 2，变量的重声明是什么意思？
	// var name = "zhouyi"
	// name, sex := "zy", "male"
	// 变量重声明的条件
	// 1，类型一致
	// 2，同一个代码块
	// 3，使用短变量声明
	// 4，被声明的变量必须是多个，并且必须有一个是新变量

	// 高内聚，低耦合（不同代码块中变量名称是可以重复的）
	// 一个程序实体的作用域总是会被限制在某个代码块中，而这个作用域最大的用处，就是对程序实体的访问权限的控制
	var block = 1
	{
		block := "inner"
		fmt.Println(block)
	}
	fmt.Println(block)

	// 不同代码块中的重名变量与变量重声明中的变量有什么区别？

	// 3，如何判断一个变量的类型？

}
