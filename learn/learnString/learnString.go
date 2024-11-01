package main

import (
	"fmt"
	"unicode/utf8"
	"strings"
)

// 字符串是不可变字节序列: 只读字节序列(string底层是字节序列)
// 字符串的默认值不是nil，而是""

// 当我看到一个字符串，我应该如何理解？
// 对于人来说，就是一堆Unicode字符（这只是我认为的数据，打印输出时，golang会找出每个字符）
// 对于计算机来说，内存或存储都默认以UTF-8的编码格式存储，就是字节序列（这TM才是实际上的数据）
// 字面量无所谓：允许十六进制、八进制和UTF编码格式

// rune类型
// 底层表达是Unicode码点
// 底层存储是UTF-8编码

// type stringStruct struct {
// 	str unsafe.Pointer
// 	len int
// }

// len(s): 返回字节数

// 为了证明什么类型都能打印
type User struct {
	name string
	age  int
}

// func (u User) String() string {
// 	return fmt.Sprintf("User{%s,%d}", u.name, u.age)
// }

type C int

func (c C) String() string {
	return fmt.Sprintf("%d", c)
}

func main() {

	// 数字
	// 任何类型都有字符串表示形式
	var a int
	a = 0x61 // a = 97  要搞清楚字面量就是一种表示，书写格式，实际类型就是int，底层存储都是一样的
	fmt.Println(a) // 数字默认打印十进制，并且结尾会加换行符
	fmt.Printf("二进制: %b\n八进制: %o\n十进制: %d\n十六进制: %x\n默认格式: %v\n", a, a, a, a)
	fmt.Printf("%T\n", a)

	var c C = 666 // golang中是不是所有类型都有默认的字符串表示形式
	fmt.Println(c) 

	// 字符串
	// 下面几种不过都是字符的书写方式
	// 中文字符
	// \u: unicode码，使用2个字节表示一个字符，后面一般跟4个十六进制的数字（一个十六进制的数字是4位，所以是16个bit，就是2个字节）
	// \x: 16进制，后面跟2位，表示一个单字节编码的字符
	// \142: ASCII码的值

	// %q: 会给字符串的值加上引号(帮助你更清晰地看到字符串在内存中的实际表示)
	// %x: 16进制
	s := "中文\x61\142\u0041"
	fmt.Printf("%T\n", s)		   // string
	fmt.Printf("%s\n", s)          // 中文abA
	fmt.Printf("%x\n", s)          // e4b8ade69687616241
	fmt.Printf("%q\n", s)          // "中文abA"
	fmt.Printf("%q\n", []rune(s))  // ['中' '文' 'a' 'b' 'A']
	fmt.Printf("%x\n", []rune(s))  // [4e2d 6587 61 62 41]      中文是3个字节，英文是1个字节(2个16进制的数字表示一个字符)
	fmt.Printf("% x\n", []byte(s)) // e4 b8 ad e6 96 87 61 62 41  中文是3个字节？？？
	fmt.Printf("%q\n", []byte(s))  // "中文abA"

	// 结构体或数组
	user := User{
		name: "zhouyi",
		age: 30,
	}
	// 通用（不通类型有不同的符号）
	// %v 输出变量的值: {zhouyi 30}       会根据不同类型，输出变量的默认格式，int类型为啥是missing
	// %+v: {name:zhouyi age:30}
	// %#v: main.User{name:"zhouyi", age:30}
	// %T: main.User
	// 如果使用%s，{zhouyi %!s(int=30)}，age字段无法格式化，说明结构体是按字段来格式化的
	fmt.Println(user)
	fmt.Printf("%v\n", user)
	fmt.Printf("%T\n", user)

	var nums []int = []int{1,2,3,4,5}
	fmt.Printf("%v\n", nums)
	fmt.Printf("%T\n", nums)

	// Bool
	var yes bool = true
	fmt.Printf("%t\n", yes)
	fmt.Printf("%v\n", yes)

	// 使用for遍历字符串，分byte和rune两种方式
	for i := 0; i < len(s); i++ { // byte
		fmt.Printf("%d: [%c];", i, s[i])
	}
	fmt.Printf("\n")

	// i: 是字节对应的索引值，不是Unicode字符对应的索引值(可以很清楚地看出每个字符是几个字节)
	for i, c := range s { // rune: 返回数组索引号，以及Unicode字符
		fmt.Printf("%d: [%c];", i, c)
	}
	fmt.Printf("\n")

	// 获取字符串中的字符数
	fmt.Printf("%d\n", utf8.RuneCount([]byte(s)))

	// 字符转换: 必须先转换成可变类型，完成后再转换回来


	// strings.Builder（只允许拼接或重置，不允许修改，当然[]byte是可以修改的，只是字符串做了限制，不允许我们修改）
	// 拼接
	s2 := "i am shaw"
	s3 := s2 + "ok" // 会把所有字符串依次拷贝到足够大的内存空间
	fmt.Printf("s3 ===>%s\n",s3) 

	var builder strings.Builder
	builder.Grow(10)
	fmt.Printf("builder length is: %d\n", builder.Len())

	// 裁剪
	fmt.Printf("s2 ===> %s\n",s2[0:3]) // 底层数组做切片
	



	// s := "My name is zhouyi\n"
	// a := []byte(s)
	// fmt.Printf("%T ,%+v\n", s, s)
	// fmt.Printf("%T ,%+v\n", a, a)

	// raw-string: 原始字符串，反斜杠没有特殊含义
	// o := `My name is zhousiyu\n`
	// fmt.Printf("%T ,%+v\n", o, o)

	// // 字节切片
	x := []byte{'x', 'y'}
	// 类型转换，string转换成[]byte
	y := []byte("zhouyi")
	fmt.Printf("%T, %s, %+v\n", x, x, x)
	fmt.Printf("%T, %s, %+v\n", y, y, y)

	// test := "周燚"
	// // 字符串切片
	// fmt.Println(test[0:3])
	// // 字节切片
	// fmt.Println([]byte(test[0:3]))

	// greeting := "My name is "
	// name := "zhouyi"

	// // 将多个字符串拼接成一个
	// z := greeting + name
	// fmt.Printf("%T ,%+v\n", z, z)

	// // 使用复合赋值运算符
	// // s := "Can you hear me?"
	// // s += "\nHear me screamin?"

	// // 将其他类型转换成字符串
	// i := 1
	// intToString := strconv.Itoa(i)
	// fmt.Println(intToString)

	// // 使用缓冲区拼接字符串
	// // 相比于+和+=，效率更高
	// var buffer bytes.Buffer
	// for i := 0; i < 500; i++ {
	// 	buffer.WriteString("z")
	// }
	// fmt.Println(buffer.String())

	// // 如何处理字符串？
	// // 将字符串转换成小写
	// fmt.Println(strings.ToLower(s))
	// // 在字符串中查找子串
	// fmt.Println(strings.Index("zhouyi", "yi"))
	// // 删除字符串中的空格
	// fmt.Println(strings.TrimSpace(" Here we go "))

}
