package main

import (
	"errors"
	"fmt"
)

// 处理错误
// 错误字符串不应该以大写字母开头或者标点符号结尾

// 理解错误类型
// 错误是一个接口类型
// type error interface {
// 	Error() string
// }

func main() {

	// 如何创建error？
	// 使用errors包中的New创建一个错误
	err := errors.New("something went wrong")

	if err != nil {
		fmt.Println(err)
	}

	// 使用fmt包设置错误字符串的格式
	name, age := "zhouyi", 25
	err2 := fmt.Errorf("%v%v", name, age)
	if err2 != nil {
		fmt.Println(err)
	}

}

// 约定：在调用可能出现问题的方法或者函数时。返回一个类型为错误的值

// 从函数返回错误
func Half(numberToHalf int) (int, error) {
	if numberToHalf%2 != 0 {
		return -1, fmt.Errorf("can not half %v", numberToHalf)
	}
	return numberToHalf / 2, nil
}

// 慎用panic
// 1， 程序处于无法恢复的状态
// 2， 发生了无法处理的错误

// 其他语言
// try-catch-finally
