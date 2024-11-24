# 并发
## 并发控制
1. channel
2. WaitGroup
    ```
    wg := sync.WaitGroup
    wg.Add(1) // 添加锁

    go func(){
        defer wg.Done // 释放锁
        do something
    }()
    wg.Wait()  // 阻塞
    ```xxx
3. Context