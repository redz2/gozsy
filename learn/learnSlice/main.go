package main

import "fmt"

// 费曼学习法: 输出是最好的输入

// 类型推断

// 变量重声明

// 可重名变量

// 类型断言

// 类型转换

// 别名类型

// 潜在类型

// 数组(array)
// 长度固定，值类型

func main() {
	// 底层数组的变化是怎样的？？？
	array := [4]int{10, 20, 30, 40}
	// {10, 20}
	// 切片容量为4，长度为2，指向array
	slice := array[0:2]
	// {10, 20, 50, 60, 70}
	// 切片容量为8，长度为5， 指向newArray
	newSlice := append(append(append(slice, 50), 60), 70)
	// 新旧slice指向的数组不一样了
	newSlice[1] += 1
	fmt.Println(array)
	fmt.Println(slice)
	fmt.Println(newSlice)
	fmt.Printf("len: %v\ncap: %v\n", len(newSlice), cap(newSlice))

	x := make([]int, 5)
	// append函数
	x = append(append(x, 1, 2, 3, 4, 5), 1)
	fmt.Println(x)
	fmt.Printf("%v %v", cap(x), len(x))
}
