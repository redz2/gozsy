# 错误处理
1. error是一个接口类型
    * 实际类型不同
    ```
    type error interface {
        Error() string
    }
    ```
2. 如果error不为nil
    * 自己设计package，设计好对外接口，而不是把具体的实现细节暴露出去
    * 我如何实现error接口
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
3. fmt.Errorf
    * 先调用fmt.Sprintf
    * 再调用errors.New
4. 如何判断错误值具体代表什么？
    * 类型已知，断言或switch
    ```
    func underlyingError(err error) error {
        switch err := err.(type) {
        case *os.PathError:
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
    ```
    * 
5. 类型断言
    ```
    t, ok := i.(T)
    ```
6. 类型选择
    ```
switch v := i.(type) {
case T:
    // v 的类型为 T
case S:
    // v 的类型为 S
default:
    // 没有匹配，v 与 i 的类型相同
}
    ```
7. 
```
type net.Error struct {
    err error
    Timeout 
    Temporary
}
```
