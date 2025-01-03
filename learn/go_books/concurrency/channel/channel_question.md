# 关于channel的几个问题
1. 通道有哪些基本特性
    * 同一个通道，发送（接收）操作之间是互斥的
    * 发送操作和接收操作都是原子操作
        * 发送到通道: 创建副本，复制到通道
        * 从通道接收: 创建通道中元素的副本，删除通道中的元素
    * 发送操作和接收操作的完整流程
        * 发送操作包括了“复制元素值”和“放置副本到通道内部”这两个步骤
        * 在这两个步骤完全完成之前，发起这个发送操作的那句代码会一直阻塞在那里，之后的代码不会执行
        
2. 通道何时会阻塞
    * 缓冲通道
        * channel已满，所有发送操作都会阻塞(还存在一个等待队列，go会优先通知先阻塞的)
        * channel为空，所有接收操作都会阻塞
        * 缓冲通道是用异步的方式进行数据传递
            * 一般情况，发送方复制到channel，再从channel复制到接收方
            * 当发送操作发送空的channel中正好有等待的接收操作，会直接把值复制给接收方
    * 非换缓冲通道
        * 无论是发送操作还是接收操作，一执行就会阻塞，直到配对操作也开始执行，才会继续传递
        * 数据是直接从发送方复制到接收方的，不会利用非缓冲通道中转
    * 值为nil的通道
        * 发送操作和接收操作都会永久处于阻塞状态，所属的goroutine中的代码永远不会执行
        * var ch2 chan int // 只声明，未初始化

3. 通道何时会panic
    * 已初始化，未关闭的channel，一定不会引起panic
    * 已关闭的channel，继续发送会引起panic
        * 通过接收表达式的第二个变量来判断通道是否关闭，有延迟(即使通道关闭，若还有值未取出，第二个变量依然为true)
        * 让发送方来关闭通道，不要让接收方来关闭
        * 接收方可以接收（零值），不会阻塞
    * 试图关闭一个已经关闭的channel，也会引起panic