package main

import (
	"fmt"
	"time"
)

var container = []string{"zero", "one", "two"}

func main() {
	container := map[int]string{0: "zero", 1: "one", 2: "two"}
	// 标准写法
	value, ok := interface{}(container).(map[int]string)
	println(value[0], ok)
	fmt.Printf("The element is %q.\n", container[1])

	type x struct {
		a string
		t time.Time
	}

	y := x{
		a: "zhouyi",
	}

	fmt.Printf("%v", y.t)
}
