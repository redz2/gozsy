package main

import "fmt"

func main() {
	// chan：表示通道类型的关键字
	// int：表示该通道类型的元素类型
	// 3：表示通道的容量（缓冲通道，非缓冲通道）
	ch1 := make(chan int, 3)
	// 2将被发送到通道
	ch1 <- 2
	ch1 <- 1
	ch1 <- 3
	// 要从该通道接收一个值
	elem1 := <-ch1
	fmt.Printf("The first element received from channel ch1: %v\n",
		elem1)
}

// 通道：一个先进先出的队列
// 可以利用通道在多个 goroutine 之间传递数据。
// 并发安全

// 对通道的发送和接收操作有哪些基本特性？

// 1，对于同一个通道，发送操作之间是互斥的，接收操作之间也是互斥的
// 对于同一个元素来说，发送和接收操作也是互斥的
// 元素值进入通道时会被复制，元素值从通道进入外界会被移动（生成副本给接收方，删除在通道中的元素值）

// 2，发送操作和接收操作中对元素值的处理都是不可分割的

// 3，发送操作在完全完成之前会被阻塞。接收操作也是如此。
// 阻塞：实现操作的互斥和元素值的完整

// 问题 1：发送操作和接收操作在什么时候可能被长时间的阻塞？