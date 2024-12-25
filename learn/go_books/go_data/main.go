package main

import (
	"fmt"
	"os"
)

func Loop3(yield func(int) bool) {
	// yield 是隐式传入的函数，不需要自己实现
	for i := range 3 {
		if !yield(i) { // 抛出值
			return
		}
	}
}

// func TrimSpace(s []byte) []byte {
// 	b := s[:0]
// 	for _, x := range s {
// 		if x != ' ' {
// 			b = append(b, x)
// 		}
// 	}
// 	return b
// }

func main() {
	s := []int{1, 2, 3}
	b := s[:1]
	b = append(b, 4)
	fmt.Println(s)
	fmt.Println(b)
	fmt.Fprintf(os.Stderr, "%s\n", "x")
	for i := range Loop3 { // 接收值
		fmt.Println(i)
	}
}
