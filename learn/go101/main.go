package main

import (
	"log"
	"runtime"
)

func main() {
	log.Println(runtime.NumCPU()) // 逻辑CPU
}
