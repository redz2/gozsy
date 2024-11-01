package main

import (
	"fmt"
	"io"
	"strings"
)

type Pet interface {
	Name()string
}

type Dog struct {
	name string
	age int
}

func (d Dog)Name()string{
	return d.name
}

func main() {
	// 接口类型: iface
	// 一个指向类型信息的指针：动态值的类型，实现接口的方法，调用方法的途径等等
	// 一个指向动态值的指针

	// 使用字面量初始化
	d := Dog{
		name: "wangwang",
		age: 23,
	}
	// 结构体每个字段都初始化为该类型的零值
	var d2 *Dog
	// 接口类型的变量不能初始化
	var p Pet
	var p2 Pet
	// 接口类型的变量的默认值是nil
	fmt.Printf("Pet1: %v,%T\n", p, p)
	fmt.Printf("Pet2: %v,%T\n", p2, p2)
	p = d
	// 我们把nil赋值给了p2，结果p2却不是nil，为啥呢？
	p2 = d2
	fmt.Printf("%v\n", d2 == nil)
	fmt.Printf("%v\n", p2 == nil)
	// 输出接口的动态类型，动态值
	// 可以输出接口类型的变量的静态类型吗？
	fmt.Printf("Pet1: %v,%T\n", p, p)
	fmt.Printf("Pet2: %v,%T\n", p2, p2) // 输出的动态类型的值nil，接口类型的值可不是nil，打印的不是自己的值


	// 创建一个strings.Reader，实现了io.Reader接口
	s1 := strings.NewReader(`zhouyi`)
	stringReader := learnInterface(s1)
	res, _ := io.ReadAll(stringReader)
	fmt.Printf("%v\n", string(res))

}

// 函数可以使用接口类型，并且返回接口类型
func learnInterface(ir io.Reader) io.Reader {
	return ir
}
