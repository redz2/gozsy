package main

import (
	"fmt"
)

func main() {
	// 结构体是值类型

	// 声明结构体
	// 相当于创建了一个模板
	type Address struct {
		Number int
		City   string
	}
	type Man struct {
		name    string
		age     int
		address Address
	}

	ad := Address{
		Number: 1,
		City:   "changzhou",
	}

	// 创建一个结构体实例，将各个数据字段设置为相应数据类型的零值 (字符串的零值是空字符串，不是nil)
	var m Man
	// {name: age:0 address:{Number:0 City:}}
	fmt.Printf("%+v\n", m)
	m.name = "zhouyi"
	m.age = 30
	m.address = ad
	fmt.Printf("%+v\n", m)

	// 使用new来创建结构体实例（为其分配内存）
	n := new(Man)
	fmt.Printf("%+v\n", n)

	// 使用简短变量赋值来创建结构体实例
	g := Man{
		name:    "zhouyi",
		age:     30,
		address: ad,
	}
	fmt.Printf("%+v\n", g)

	h := g
	fmt.Printf("%+v\n", h)

	// 结构体嵌套有什么用？

	// 比较结构体
	// 使用reflect.TypeOf(m)

	// 创建一个匿名结构体，一般不用
	// m := struct {
	// 	name string
	// 	age  int
	// }{
	// 	name: "zhouyi",
	// 	age:  30,
	// }
	// fmt.Printf("%v\n", m)

	// 方法调用是否会修改值？
	// 主要看方法的接收者是值类型还是指针类型
	c1 := Child{age: 1} // 值类型
	fmt.Printf("%v", c1)
	c1.GrowUp()
	fmt.Printf("%v", c1)
	c1.TryGrowUp()
	fmt.Printf("%v", c1)

	fmt.Printf("\n-----------------------\n")

	c2 := &Child{age: 1} // 指针类型
	fmt.Printf("%v", c2)
	// 指针接收者的方法会修改调用者
	// 无论调用者是值类型还是指针类型，都会强制转换成指针类型，age都会发生改变
	// c1: 值类型，实际(&c1).GrowUp()
	// c2: 指针类型
	c2.GrowUp()
	fmt.Printf("%v", c2)

	// 值接收者的方法不会修改调用者
	// 无论调用者是值类型还是指针类型，都会强制转换成值类型，age都不会改变
	// c1: 值类型，直接复制
	// c2: 指针类型，等价于 (*c2).TryGrowUp()
	c2.TryGrowUp()
	fmt.Printf("%v", c2)

	fmt.Printf("\n-----------------------\n")

	// 如何给接口赋值？
	// 1，Child结构体实现了People接口
	// 2，如果有方法的接收者是指针类型， 那么只能用指针类型给接口赋值
	// 3，因为指针类型对数据结构是读写权限，不能缩小权限！！！值类型是读权限（虽然调用方法会强制转换）
	var p1 People = &Child{age: 1} // p1是一个接口
	fmt.Printf("%v", p1)
	c1.TryGrowUp()
	fmt.Printf("%v", p1)
	p1.GrowUp()
	fmt.Printf("%v", p1)
}

type People interface {
	Age() int
	GrowUp()
}

type Child struct {
	name string
	age  int
}

func (c Child) Age() int {
	return c.age
}

func (c *Child) GrowUp() {
	c.age++
}

func (c Child) TryGrowUp() {
	c.age++
}
