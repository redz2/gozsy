package learn

import (
	"fmt"
	"reflect"
	"strconv"
)

func sayHello(s string) string {
	return "Hello " + s
}

func test() {
	fmt.Println((sayHello("zhouyi")))

	// 布尔类型
	// 如果没有给布尔变量赋值，默认为false
	var b bool
	fmt.Println(b)

	// 数值类型
	var i int = 3
	fmt.Println(i)

	// 浮点数
	var f float64 = 0.111
	fmt.Println(f)

	// 检查变量的类型
	fmt.Println(reflect.TypeOf(i))
	fmt.Println(reflect.TypeOf(f))
	fmt.Println(reflect.TypeOf(b))

	// 类型转换
	s1 := strconv.FormatBool(true)
	// string类型的字符串，十进制，大小限制，返回int64
	s2, _ := strconv.ParseInt("12345678", 10, 64)
	// string类型的字符串，大小限制，返回float64
	s3, _ := strconv.ParseFloat("123.12312", 64)
	fmt.Println(s1, s2, s3)
}
