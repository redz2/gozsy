# channel
1. Don’t communicate by sharing memory; share memory by communicating
2. channel通道（先进先出队列）
    * 基本特性
        * 对于同一个通道，发送操作之间是互斥的，接收操作之间也是
            * 同一时刻，只会执行对同一个通道的任意发送(或接收)操作的其中一个，直到该操作完成
        * 发送操作和接收操作对元素值的处理是不可分割的(原子性)
            * 发送到通道: 创建副本，复制到通道
            * 从通道接收: 创建通道中元素的副本，删除通道中的元素
        * 发送操作在完全完成之前会被阻塞，接收操作也是
            * 发送操作包括了“复制元素值”和“放置副本到通道内部”这两个步骤
            * 在这两个步骤完全完成之前，发起这个发送操作的那句代码会一直阻塞在那里，之后的代码不会执行
            * 发送操作完成后，go会通知goroutine，让他
    ```
    // 声明并初始化一个通道
    ch1 := make(chan int, 3)
    // 向通道发送一个值
    ch1 <- 2    // 该操作是原子操作，并且多个发送操作互斥
    ch1 <- 1
    ch1 <- 3
    // 从该通道接收一个元素值
    elem1 := <-ch1    // 该操作是原子操作，并且多个接收操作互斥，如果channel中没有值，会阻塞在这里
    fmt.Printf("The first element received from channel ch1: %v\n", elem1)
    ```
3. channel何时会阻塞？
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
4. channel何时会panic？
    * 已初始化，未关闭的channel，一定不会引起panic
    * 如果channel被关闭，发送操作会引起panic
        * 接收操作会感知到channel被关闭
        * 通过接收表达式的第二个变量来判断通道是否关闭，有延迟(即使通道关闭，若还有值未取出，第二个变量依然未true)
        * 让发送方来关闭通道，不要让接收方来关闭
    * 试图关闭一个已经关闭的channel，也会引起panic
5. 单向通道
    * var chan <- int 发送到channel，发送通道
    * var <- chan int 从channel接收，接收通道
    * 通道是为了传递数据的，单向通道一般是没啥用的
    * 单向通道有啥用？
        * 约束其他代码的行为，下面的函数ch是一个发送通道，只能发送值到通道中，无法从通道中接收值
        * 调用函数时，可以传一个双向通道，go会自动转换成单向通道
        ```
        func SendInt(ch chan <- int){
            ch <- rand.Intn(1000) // 只能发送值到ch中
        }
        ```
        * 接口声明中，用来约束实现
        ```
        type Notifier interface {
            SendInt(ch chan <- int)
        }
        ```
        * 返回一个channel
        ```
        func getIntChan() <-chan int {
            num := 5
            ch := make(chan int, num)
            for i := 0; i < num; i++ {
                ch <- i
            }
            close(ch)
            return ch
        }

        intChan2 := getIntChan()
        for elem := range intChan2 {                    // for会尝试从通道中取值，如果没有值会阻塞(如果是channel为nil，for语句会永远阻塞)
            fmt.Printf("The element in intChan2: %v\n", elem)
        }
        ```
6. select(channel专用)
    * 如果没有default分支，并且表达式都阻塞，select就会被阻塞
    * 如果有default分支，即使表达式都阻塞，select也不会阻塞
    * 如果多个分支就绪，执行哪一个分支呢？
    ```
    select (
        case <- ch1:
            do something
        case <- ch2:
            do something
        default:
            do something
    )
    ```

    ```
    intChan := make(chan int, 1)
    // 一秒后关闭通道。
    time.AfterFunc(time.Second, func() {
        close(intChan)
    })
    select {
    // 一开始是阻塞的
    case _, ok := <-intChan:
        if !ok {
            fmt.Println("The candidate case is closed.")
            // 如果通道关闭就退出select
            break
        }
        fmt.Println("The candidate case is selected.")
    }
    ```