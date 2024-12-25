# channel
1. 通道介绍
    * 通道的主要作用是实现并发同步
        * 解决协程之间通信问题
    * 不要让计算通过共享内存来通讯，而应该让它们通过通讯来共享内存（一个计算就是一个协程）
        * 共享内存: 互斥锁
        * 通道: 一些协程可以向通道发送数据，一些协程可以从通道获取数据（随着数据的发送和接收，数据的所有权在协程之间转移了）

2. 通道的类型和值
    * 按数据传递方向分类
        * 双向通道: chan T（可以隐式转换成单向通道，反之不行，显式转换也不行）
        * 单向通道
            * 发送通道: chan<- T
            * 接收通道: <-chan T
        * 单向通道有啥用？
            * 通道是为了传递数据，单向通道一般是没啥用的
            * 约束其他代码的行为，下面的函数ch是一个发送通道，只能发送值到通道中，无法从通道中接收值（控制读写权限）

    * 按通道容量分类
        * 非缓冲通道: 容量为0
        * 缓冲通道

    * 通道的零值: nil
        * 非零通道的创建: make(chan int, 10)

3. channel的3种操作（读、写、关闭）
    ```
    // 关闭一个channel
    close(ch) 

    // 向通道发送一个值
    ch <- v

    // 从通道接收一个值（如果channel已关闭，会读到零值，但是可以一直读取，不会panic，可以通过第二个返回值判断通道是否关闭）
    v = <-ch

    // 查询一个通道的容量
    cap(ch)

    // 查询一个通道的长度(发送到此通道还未被接收的值的数量)
    len(ch)
    ```

4. 通道操作详解
    * channel的3种状态
        * nil通道
        * 非nil通道但已关闭（关闭的channel的值可不是nil）
        * 非nil通道但未关闭
    * 一些操作结果
        * 关闭一个nil通道或者已关闭的通道将panic
        * 向一个已关闭通道发送数据将panic
        * 向一个nil通道发送或接收数据将永久阻塞
    * 原理
        * 接收数据的协程队列（处于阻塞状态，等待接收数据）
        * 发送数据的协程队列（处于阻塞状态，等待发送数据）
        * 数据缓冲队列
        * 每个通道内有一个互斥锁
    * 使用for-range读取channel（channel关闭会自动退出）
    * 使用 _,ok 判断channel是否关闭
    * 使用select处理多个channel

5. 通道Example
* 数据传递、信号传递
    ```
    c := make(chan int)
    go func(ch chan<- int, x int){
        time.Sleep(time.Second)
        ch <- x // 生产者
    }(c, 3) 

    done := make(chan struct{}) // 结束信号
    go func(ch <-chan int){
        x := <-ch // 消费者
        time.Sleep(time.Second)
        done <- struct{}{}
    }(c)
    <-done // 阻塞主线程
    ```

* 将通道用作future/promise
    * 异步编程（一个协程就是一个任务，通过channel等待协程完成，进行后续处理，比如结果汇总）
    ```
    func longTimeRequest() <-chan int32 {
        r := make(chan int32)

        go func() {
            time.Sleep(time.Second * 3) // 模拟一个工作负载
            r <- rand.Int31n(100)
        }()

        return r
    }

    func sumSquares(a, b int32) int32 {
        return a*a + b*b
    }

    func main() {
        rand.Seed(time.Now().UnixNano()) // Go 1.20之前需要

        a, b := longTimeRequest(), longTimeRequest()
        fmt.Println(sumSquares(<-a, <-b))  // 是不是很像await
    }
    ```

    ```
    func longTimeRequest(r chan<- int32)  {
        time.Sleep(time.Second * 3) // 模拟一个工作负载
        r <- rand.Int31n(100) // 生产者
    }

    func sumSquares(a, b int32) int32 {
        return a*a + b*b
    }

    func main() {
        rand.Seed(time.Now().UnixNano()) // Go 1.20之前需要

        ra, rb := make(chan int32), make(chan int32)
        go longTimeRequest(ra)
        go longTimeRequest(rb)

        fmt.Println(sumSquares(<-ra, <-rb))
    }
    ```

    * 多个数据源返回相同的数据，使用返回最快的那一份数据
    ```
    func source(c chan<- int32) {
        ra, rb := rand.Int31(), rand.Intn(3) + 1
        // 睡眠1秒/2秒/3秒
        time.Sleep(time.Duration(rb) * time.Second)
        c <- ra
    }

    func main() {
        rand.Seed(time.Now().UnixNano()) // Go 1.20之前需要

        startTime := time.Now()
        c := make(chan int32, 5) // 必须用一个缓冲通道
        for i := 0; i < cap(c); i++ {
            go source(c)
        }
        rnd := <- c // 只有第一个回应被使用了
        fmt.Println(time.Since(startTime))
        fmt.Println(rnd)
    }
    ```

* 使用通道实现消息通知
    * 不关心值是什么，只关心回应是否发生
    * make(chan struct{})    // 一般使用空结构体，尺寸为0
    * 单对单
    ```
    values := make([]byte, 32 * 1024 * 1024)
	if _, err := rand.Read(values); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	done := make(chan struct{}) // 也可以是缓冲的

	// 排序协程
	go func() {
		sort.Slice(values, func(i, j int) bool {
			return values[i] < values[j]
		})
		done <- struct{}{} // 通知排序已完成
	}()


	<- done // 等待通知
	fmt.Println(values[0], values[len(values)-1])
    ```

    ```
    done := make(chan struct{})

	go func() {
		fmt.Print("Hello")
		time.Sleep(time.Second * 2)

		<- done 
	}()

	done <- struct{}{} // 阻塞在此，等待通知（上面的写法更好理解）
	fmt.Println(" world!")
    ```

    * 多对单
    * 定时通知（timer）
    ```
    func AfterDuration(d time.Duration) <- chan struct{} {
        c := make(chan struct{}, 1)
        go func() {
            time.Sleep(d)
            c <- struct{}{}
        }()
        return c
    }

    func main() {
        fmt.Println("Hi!")
        <- AfterDuration(time.Second)
        fmt.Println("Hello!")
        <- AfterDuration(time.Second)
        fmt.Println("Bye!")
    }
    ```

* 将通道用作互斥锁（mutex）
    * 通过发送操作来加锁
    * 通过接收操作来解锁
    ```
    m := make(chan struct{}, 1) // 容量必须为1
    counter := 0
    increase := func(){
        m <- struct{}{}  // 通过发送操作来加锁（如果通过接收操作来加锁，必须提前放一个锁）
        counter++        // 保证数据一致性
        <-m // 解锁
    }

    increase1000 := func(done chan struct{}){
        for i:=0; i<1000; i++ {
            increase()
        }
        done <- struct{}{}
    }

    done := make(chan struct{})
    go increase1000(done)
    go increase1000(done)
    <-done; <-done

    ```

* 将通道用作计数信号量
    * 缓冲通道容量是N，表示最多有N个协程可以执行（限制最大并发量）

* 对话（乒乓）
    * 两个协程通过一个channel进行通话

* 检查通道的长度和容量
* 使当前协程永久阻塞

6. select: channel专用
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
