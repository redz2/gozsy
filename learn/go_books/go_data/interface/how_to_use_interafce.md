# 接口的使用
1. 尽量定义小接口（接口可以组合使用）
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

2. 尽量避免使用空接口作为函数参数
    * 空接口不提供信息，只能通过反射来获取信息（然而反射性能差）
    * 标准库中有哪些函数这么干了？
        * 容器算法: sort
        * 格式化: fmt,log

3. 接口值的比较
    * 和nil可以比较
    * 接口之间比较可能有风险

4. 使用接口作为程序水平组合的连接点
    * 垂直组合
        * 通过嵌入接口构建接口
        * 通过嵌入接口构建结构体
        * 通过嵌入结构体构建结构体

    * 水平组合
        * __接口到底是什么？__(或者说到底包裹了啥)
            * 之前我认为接口能包含不同数据类型，md还有函数类型
                * 数据类型 -> 包裹不同数据类型
                * 函数类型 -> 包裹不同处理逻辑
            * 接口是函数，还是数据？弄清楚这一点很重要？

        * 包裹数据类型，那就实现了 __多态__: func FuncName(p YourInterfaceType)
        ```
        // 函数参数是接口类型，该函数可以处理不同类型的数据
        func readAll(r io.Reader, capacity int64)(b []byte, err error)
        ```

        * 输入接口类型，输出接口类型: func FuncName(p YourInterfaceType) YourInterfaceType
            * 如果包裹的是函数，类似于装饰器函数的功能
            * 如果包裹的是数据，类似于普通函数

        * 装饰器函数: func Fn(p FuncType) FuncType（__目前和接口无关__）
            * 作用: 扩充函数功能
            ```
            type AddFunc func(int, int)int
            func Wrap(f AddFunc) AddFunc {
                return func(a,b int){
                    return f(a,b) + 1000
                }
            }
            ```
            * 泛型的装饰器: 因为装饰器函数的参数中耦合了被装饰函数的类型，所以装饰器函数无法做到通用（比如计算函数执行时间）
                * python的装饰器就比较变态了，可以用在所有函数上

        * 适配器函数类型（使用接口包裹不同的函数类型，作用就是约束函数类型）__接口包裹函数类型__
            * 实际上ListernAndServe只需要传入一个函数就行了，为啥非得要接口类型
                * 通过接口来约束函数类型

            * 目前我们需要实现Handler接口
                * 需要定义一个数据类型
                    * 其实不需要任何数据，只需要处理逻辑
                    * 解决办法: 定义一个函数类型(适配器函数类型)，让这个函数类型实现该接口，我们只需要定义函数了
                        * 这个接口中不需要数据，只需要处理逻辑时，可以定义一个函数类型来实现该接口
                        * 而不是去定义数据类型
                * 实现接口中的ServeHTTP方法
            ```
            // 我们的目标是实现Handler接口
            type Handler interface {
                ServeHTTP(ResponseWriter, *Request)
            }

            // 我们定义一个函数类型，并让这个函数类型实现该接口
            // 我们其实是不需要数据的，只需要处理逻辑
            type HandlerFunc func(ResponseWrite, *Request)

            func (f HandlerFunc) ServeHttp(w ResponseWriter, r *Request){
                f(w, r)
            }

            // 至此，该函数类型实现了Handler接口
            // 函数类型实现接口 vs 数据类型实现接口
            // 目前能想到的优势: 可以添加任意处理逻辑

            func greetings(w ResponseWrite, r *Request){
                ...
            }

            // Handler接口包裹的其实是不同的处理逻辑，可以认为就是一个函数
            // 那么为什么不直接写个函数呢？非要写接口类型？可以做函数约束
            http.ListernAndServe(":8080", http.Handler(http.HandlerFunc(greetings))) 
                                                        // greetings -> http.HandlerFunc -> http.Handler（接口可以隐式转换）
            ```

        * 中间件: func Fn(h Handler)Handler
            * Handler中包裹了函数，这tm不就是个装饰器函数吗？
        ```
        func Logging(h http.Handler) http.Handler {

            // return a new http.Handler
            // every HandlerFunc is a http.Handler
            return http.Handler(func(w ResponseWrite, r *Request){
                // add log
                h.ServeHTTP(w, r)
            })
        }

        http.ListernAndServe(":8080", Logging(http.Handler(greetings)))
        ```

        * 如果有多个中间件，怎么添加中间件？多个修饰器的 Pipeline
        ```
        // Middle是中间件函数
        // Handler是处理函数
        type Middle func(h http.Handler) http.Handler

        // 中间件工厂函数
        func Logging() Middle {
            return func(h http.Handler) http.Handler {
                return http.Handler(func(w ResponseWrite, r *Request){
                    // add log
                    h.ServeHTTP(w, r)
                })
            }
        }

        // 每个Middle会去装饰Handler
        func Chain(h http.Handler, middles ..Middle){
            for _, m := range middle {
                f = m(f)
            }
            return f
        }
        http.ListernAndServe(":8080", chain(greetings, Logging()))
        ```