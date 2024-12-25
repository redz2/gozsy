package main

import (
	"fmt"
	"time"
)

func init() {
	// 程序开始运行就会记录一个时间戳
	time.Sleep(1 * time.Second)
}

func main() {
	// 定时器
	// 1. time.Sleep: 阻塞Goroutine
	time.Sleep(1 * time.Second)
	fmt.Println(time.Now())
	time.Sleep(3 * time.Second)
	fmt.Println(time.Now())
	// 2. time.Timer
	timer := time.NewTimer(time.Second)
	fmt.Println(<-timer.C)
	timer.Reset(3 * time.Second)
	fmt.Println(<-timer.C)
	timer.Stop()
	// 3. time.After
	fmt.Println(<-time.After(time.Second))

	fmt.Println("test ticker")
	// 使用Ticker实现周期定时
	tiker := time.NewTicker(time.Second)
	count := 0
	for {
		fmt.Println(<-tiker.C)
		count++
		if count > 10 {
			break
		}
	}
}
