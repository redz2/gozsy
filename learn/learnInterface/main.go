package main

import (
	"fmt"
	"io"
	"strings"
)

func main() {
	s1 := strings.NewReader(`zhouyi`)
	lI := learnInterface(s1)
	asd, _ := io.ReadAll(lI)
	fmt.Printf("%v\n", string(asd))
}

// 函数可以使用接口类型，并且返回接口类型
func learnInterface(ir io.Reader) io.Reader {
	return ir
}
