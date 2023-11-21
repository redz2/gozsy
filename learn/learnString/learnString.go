package main

import (
	"bytes"
	"fmt"
	"strconv"
	"strings"
)

// 理解字符串是什么？
// 字符串实际上是只读的字节切片
// len(s): 返回字节数

// ASCII码
// Unicode符号集，只规定了如何编码，没有规定如何存储
// UTF-8

func main() {

	// 解释型字符串，反斜杠有特殊含义
	s := "My name is zhouyi\n"
	a := []byte(s)
	fmt.Printf("%T ,%+v\n", s, s)
	fmt.Printf("%T ,%+v\n", a, a)

	// 原始字符串，反斜杠没有特殊含义
	o := `My name is zhousiyu\n`
	fmt.Printf("%T ,%+v\n", o, o)

	// 字节切片
	x := []byte{'x', 'y'}
	// 类型转换，string转换成[]byte
	y := []byte("zhouyi")
	fmt.Printf("%T ,%+v\n", x, x)
	fmt.Printf("%T ,%+v\n", y, y)

	test := "周燚"
	// 字符串切片
	fmt.Println(test[0:3])
	// 字节切片
	fmt.Println([]byte(test[0:3]))

	greeting := "My name is "
	name := "zhouyi"

	// 将多个字符串拼接成一个
	z := greeting + name
	fmt.Printf("%T ,%+v\n", z, z)

	// 使用复合赋值运算符
	// s := "Can you hear me?"
	// s += "\nHear me screamin?"

	// 将其他类型转换成字符串
	i := 1
	intToString := strconv.Itoa(i)
	fmt.Println(intToString)

	// 使用缓冲区拼接字符串
	// 相比于+和+=，效率更高
	var buffer bytes.Buffer
	for i := 0; i < 500; i++ {
		buffer.WriteString("z")
	}
	fmt.Println(buffer.String())

	// 如何处理字符串？
	// 将字符串转换成小写
	fmt.Println(strings.ToLower(s))
	// 在字符串中查找子串
	fmt.Println(strings.Index("zhouyi", "yi"))
	// 删除字符串中的空格
	fmt.Println(strings.TrimSpace(" Here we go "))

}
