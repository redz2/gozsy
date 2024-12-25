# 错误处理
1. error是一个接口类型
```
type error interface {
    Error() string
}
```

2. 如何创建一个实现error接口的具体类型？
    * errors.New()
    * fmt.Errorf()
        * 先调用fmt.Sprintf，再调用errors.New
    * 其他函数返回error
    ```
    func echo(request string) (response string, err error) {
        // 卫述语句，检查后续操作的前置条件
        if request == "" {
            // 静态类型是error
            // 动态类型是error包中私有的*errorString
            err = errors.New("empty request")
            return
        }
        response = fmt.Sprintf("echo: %s", request)
        return
    }
    ```
    
3. 如何判断错误值具体代表什么？
* 类型已知，断言或switch
```
func underlyingError(err error) error {
    switch err := err.(type) {  // 类型断言: t, ok := i.(T)
    case *os.PathError:         // 类型选择
        return err.Err
    case *os.LinkError:
        return err.Err
    case *os.SyscallError:
        return err.Err
    case *exec.Error:
        return err.Err
    }
    return err
}

switch v := i.(type) {
case T:
    // v 的类型为 T
case S:
    // v 的类型为 S
default:
    // 没有匹配，v 与 i 的类型相同
}
```

4. 错误处理策略（Go大部分函数代码结构都相似，先进行一系列初始检查，防止错误发生，之后是函数的实际逻辑）
    * 传播错误
    ```
    fmt.Errorf("parsing %s as HTML: %v", url, err)  // 既包含了err，又补充了新的信息，形成错误链条
    ```
    * 错误偶发，需要重新尝试
    ```
    func WaitForServer(url string) error {
        const timeout = 1 * time.Minute
        deadline := time.Now().Add(timeout)  // 超时机制
        for tries := 0; time.Now().Before(deadline); tries++ { // 重试
            _, err := http.Head(url)
            if err == nil {
                return nil // success
            }
            log.Printf("server not responding (%s);retrying…", err)
            time.Sleep(time.Second << uint(tries)) // exponential back-off
        }
        return fmt.Errorf("server %s failed to respond after %s", url, timeout)
    }
    ```
    * 错误发生后，无法继续执行
    ```
    // (In function main.)
    if err := WaitForServer(url); err != nil {
        // log默认会输出时间
        // log.Fatalf("Site is down: %v\n", err) 
        fmt.Fprintf(os.Stderr, "Site is down: %v\n", err)
        os.Exit(1)
    }
    ```
    * 有时只需要打印错误消息，不需要中断或者返回
    ```
    if err := Ping(); err != nil {
        log.Printf("ping failed: %v; networking disabled",err)
    }
    ```
    * 直接忽略错误
    ```
    os.RemoveAll(dir) // ignore errors; $TMPDIR is cleaned periodically
    ```

5. EOF
```
in := bufio.NewReader(os.Stdin)
for {
    r, _, err := in.ReadRune()
    if err == io.EOF {
        break // finished reading
    }
    if err != nil {
        return fmt.Errorf("read failed:%v", err)
    }
    // ...use r…
}
```
