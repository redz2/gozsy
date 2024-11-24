# 接口
0. 类型
    * 接口类型
        * 基本接口类型（方法集）
    * 非接口类型
1. 接口类型变量
    * nil error 不等于 nil？nil只是简单
    ```
    type MyInt interface {}
    var pa *int = nil     // 此时的nil只是一种代码表示，这里会将无类型的nil隐式转换为*int的nil 
    var myint MyInt = pa  // 只有未初始化的接口值才为nil
    if myint == nil {
        ...
    }
    ```
    * 值包裹: 接口类型的值可以用来包裹非接口类型的值
        * 什么是接口: 一个用来包裹非接口类型值的盒子
        * 接口赋值的一些细节
            * 官方说法: 如果类型T实现了接口I，T可以隐式转换成I
                * 这就是为什么函数参数使用接口，可以传入具体类型
                * 函数参数使用接口，其实就实现了多态（扩大函数的使用范围）
            * 如果T是非接口值（动态类型 + 动态值），T的复制会被包裹在结果I值中
                * 动态值的直接部分无法修改
            * 如果T也是接口值，T包裹的值会复制一份到I值中
        ```
        // 一个*Book值被包裹在了一个Aboutable值中。
        var a Aboutable = &Book{"Go语言101"}

        // i是一个空接口值。类型*Book实现了任何空接口类型。
        var i interface{} = &Book{"Rust 101"}

        // Aboutable实现了空接口类型interface{}。
        i = a
        ```
        * Go构建一个表来记录各个类型的信息（类型种类，以及此类型的方法、字段、尺寸）
        * 接口包含什么？
            * 动态类型的信息: 动态类型 + 动态值 -> 实现反射
            * 一个方法表（切片） -> 实现多态

2. 接口类型的内部表示
    * eface: 空接口
    ```
    type eface struct {
        _type *_type       // 动态类型信息
        data unsafe.Pointer
    }
    ```
    * iface: 普通接口
    ```
    type iface struct {
        tab *itab        // 动态类型信息 + 方法列表 + 非空接口信息
        data unsafe.Pointer  // 动态类型变量的值
    }
    ```
3. 尽量定义小接口
    ```
    type error interface {
        Error() string
    }

    type Reader interface {
        Read(p []byte)(n int, err error)
    }

    type Handler interface {
        ServerHttp(ResponseWrite, *Request)
    }
    ```
4. 尽量避免使用空接口作为函数参数
    * 空接口不提供信息，只能通过反射来获取信息
    * 标准库
        * 容器算法: sort
        * 格式化: fmt,log
5. 使用接口作为程序水平组合的连接点
    * 垂直组合
        * 通过嵌入接口构建接口
        * 通过嵌入接口构建结构体
        * 通过嵌入结构体构建结构体
    * 水平组合
        * func FuncName(p YourInterfaceType)
        ```
        // 处理不同类型的数据
        func readAll(r io.Reader, capacity int64)(b []byte, err error)
        ```
        * 包裹函数: func FuncName(p YourInterfaceType) YourInterfaceType
            * 形成链式调用
        ```
        通过包裹函数就能添加额外功能了

        ```
        * 适配器函数类型: func Fn(p FuncType) FuncType
            * 上面的特例，函数类型实现了该接口
            * 把实现接口变成了实现一个函数！！！
            ```
            type Handler interface {
                ServeHTTP(ResponseWriter, *Request)
            }

            // 要记住一点，我们的最开始的类型是http.Handler!!!
            type HandlerFunc func(ResponseWrite, *Request)

            func (f HandlerFunc) ServeHttp(w ResponseWriter, r *Request){
                f(w, r)
            }

            // 那么，任何HandlerFunc类型的函数，都实现了Handler接口
            // 那我们定义一个HandlerFunc，我们就能把这个函数转换成接口类型

            // 怎么就定义函数就实现接口了
            func greetings(w ResponseWrite, r *Request){
                ...
            }

            http.ListernAndServe(":8080", http.Handler(greetings))
            ```
        * 中间件
        ```
        // 包裹接口类型
        func Logging(h http.Handler) http.Handler {

            // return a new http.Handler
            // every HandlerFunc is a http.Handler
            return http.Handler(func(w ResponseWrite, r *Request){
                // add log
                // old handler need execute
                h.ServeHTTP(w, r)
            })
        }

        http.ListernAndServe(":8080", Logging(http.Handler(greetings)))
        ```
        * 如果把中间件也当成函数类型
        ```
        type Middle func(h http.Handler) http.Handler

        func Chain(h http.Handler, middles ..Middle){
            for _, m := range middle {
                f = m(f)
            }
            return f
        }
        http.ListernAndServe(":8080", chain(greetings, Logging))
        ```