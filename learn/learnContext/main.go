package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {

	// context本质是为了处理go的并发控制问题

	// 多个任务并行控制，等待所有任务完成
	// taskNum := 3
	// wg := sync.WaitGroup{}
	// wg.Add(taskNum)
	// for i := 0; i < taskNum; i++ {
	// 	go func(taskNo int) {
	// 		fmt.Printf("taskNo: %v is running\n", taskNo)
	// 		time.Sleep(4 * time.Second)
	// 		wg.Done()
	// 	}(i)

	// }
	// 阻塞在这里！！！
	// wg.Wait()

	// 控制任务的停止
	// 1，数据通道关闭

	// data := make(chan int, 10)

	// go func(data chan int) {
	// 	for {
	// 		select {
	// 					阻塞在这里！！！
	// 		case val, ok := <-data:
	// 			time.Sleep(3 * time.Second)
	// 			if !ok {
	// 				fmt.Println("channel closed")
	// 				return
	// 			}
	// 			fmt.Println(val)
	// 		}
	// 	}
	// }(data)

	// go func() {
	// 	data <- 1
	// 	data <- 2
	// 	data <- 3
	// 	data <- 4
	// 	// 即使channel已经关闭，如果channel中还有数据，依然会取完
	// 	close(data)
	// 	fmt.Println("data is closed")
	// }()

	// time.Sleep(10 * time.Second)

	// 2，单独退出通道
	data := make(chan int, 10)
	defer close(data)
	exit := make(chan struct{})
	taskNum := 3

	wg := sync.WaitGroup{}
	wg.Add(taskNum)

	// 创建3个携程用于处理任务
	for i := 0; i < taskNum; i++ {
		// 创建3个任务
		go func(taskNo int, data chan int, exit chan struct{}) {

			fmt.Printf("Now task No %v is running\n", taskNo)
			// 函数退出时执行
			defer wg.Done()

			for {
				select {
				// 阻塞在这！！！
				case val, ok := <-data:
					if !ok {
						fmt.Printf("task %v: data channel is closed\n", taskNo)
						return
					}
					fmt.Printf("task %v: value is %v\n", taskNo, val)
				// 阻塞在这！！！
				case <-exit:
					fmt.Printf("taks %v: exit channel is closed\n", taskNo)
					return

				// 超时控制
				case <-time.After(5 * time.Second):
					fmt.Printf("task %v timeout\n", taskNo)
					return
				}

			}

		}(i, data, exit)
	}

	// 等协程创建好
	time.Sleep(3 * time.Second)
	go func() {
		// 传一些值过去，控制任务的停止
		data <- 1
		data <- 2
		data <- 3
		data <- 4
		// 如果超过5s，任务就会因为超时退出
		time.Sleep(3 * time.Second)
		// 关闭exit channel
		close(exit)

	}()

	// 所有协程退出程序才能继续
	// 等待锁释放，也会阻塞
	wg.Wait()

	fmt.Println("last executed")
}

// 既然channel可以处理这些问题，为什么要引入context？
// 1，支持多级嵌套，父任务停止后，子任务自动停止
// 2，控制停止顺序，先停哪个，再停哪个

type Context interface {
	// 当 context 被取消或者到了 deadline，返回一个被关闭的 channel
	Done() <-chan struct{}

	// 在 channel Done 关闭后，返回 context 取消原因
	Err() error

	// 返回 context 是否会被取消以及自动取消时间（即 deadline）
	Deadline() (deadline time.Time, ok bool)

	// 获取 key 对应的 value
	Value(key interface{}) interface{}
}
