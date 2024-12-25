# 并发
## 并发和并行
1. 并发不是并行
    * 并行就是同时做很多事情
        * 要把任务切分成 __合适大小__ 的单元(太大太小都有问题)
    * 并发是同时管理很多事情（使用较少的资源做更多的事情）

2. Go并发编程基础: http://blog.xiayf.cn/2015/05/20/fundamentals-of-concurrent-programming/

## 并发控制
1. channel
    * 控制任务停止
    ```
    // 关闭通道
    
    // 创建一个单独的退出通道，exit
    ```
2. WaitGroup
    ```
    wg := sync.WaitGroup
    wg.Add(1) // 添加锁

    go func(){
        defer wg.Done // 释放锁
        do something
    }()

    wg.Wait()  // 阻塞
    ```
3. Context