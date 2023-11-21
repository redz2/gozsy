package main

import "fmt"

// 数组
// 长度固定，属于值类型
// 在go语言中，数组是值类型。这意味着，数组在传递时，传递的是原数组的拷贝

// 切片
// 有一个窗口，你能通过这个窗口看到一个数组，但是不能够看到数组的全部元素，只能看到连续的一部分元素
// 可变长，属于引用类型
// 长度，容量，指向数组的指针

// 切片的多种声明方式，nil切片和空切片的区别
// s := make([]byte, 5) ===> 指针不为nil，已经为其分配了底层的数组，数组可能也是空的
// s := []byte{} ===> 同上
// var s []byte ===> 指针为nil，不知道底层数组在哪里

// 扩容
// 可以把slice理解为一种”动态数组“
// 1, 如果切片容量小于1024，扩容时cap*2；如果切片容量大于1024，扩容时cap*1.25
// 2, 如果扩容后，没有触及原数组的容量，切片中指针还是指向原数组；如果超过原数组，go开辟一块新内存，把原数组的值拷贝过来

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
