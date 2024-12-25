package main

import (
	// . "log" 不建议这么用 should not use dot imports
	"fmt"
	"log"
	"runtime"
)

func main() {
	a := 1
	fmt.Printf("%T\n", a)
	log.Println(runtime.NumCPU()) // 逻辑CPU
}
