package main

import (
	"fmt"
)

type Tx interface {
	String() string
}

type Tt struct {
	name string
}

func (t Tt) String() string {
	return fmt.Sprint(t.name)
}

// 参数必须实现Tx接口
// 返回实现Tx的接口
func ReturnTt(x Tx) Tt {
	fmt.Println(x)
	z := Tt{
		name: "xx",
	}
	return z
}

func main() {
	x := Tt{
		name: "zhouyi",
	}
	// var y Tx = ReturnTt(x)
	// var z Tx = x
	fmt.Println(x.name)
	fmt.Println(ReturnTt(x).name)
}
