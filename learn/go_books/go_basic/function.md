# 函数
0. 函数式编程
    * 准则: 不依赖于外部的数据，而且也不改变外部数据的值，而是返回一个新的值给你
        * 既不要把你程序的细节暴露给外界，也尽量不要让外界的变动影响到你的程序
        * 函数中不要存在外部变量
    * 函数是什么？
        * 函数和其他值一样，拥有类型，那么函数和其他值的区别是什么？
            * 其他值表示数据，函数表示处理逻辑
            * 可以像值一样赋值给其他变量（或者作为函数参数）
                * 如果函数不是值，那么函数参数只能是值；如果函数也是值，那么函数参数还能是处理逻辑
            * 函数零值为nil，调用会panic
            * 函数值只能和nil比较，函数之间不能相互比较
        * 即使函数类型实现了某个接口，记住一点，函数表示的是处理逻辑
    ```
    // 非函数式编程
    var a int
    func Add(){
        return a + 1  // a 是全局变量
    }
    ```
    * 把函数当成变量来用，关注描述问题，而不是怎么实现（有返回值的函数可以被认为是表达式）
    * 函数中的变量是分配在堆上还是栈上
        * 如果返回的是指针，就会分配在堆上（会参与垃圾回收）
        * 如果返回的是值，就会分配在栈上（临时变量，不需要被外部使用）

1. 函数声明
    * 可以把参数和结果看作是省略了var关键字的标准变量声明
    ```
    // 函数式编程
    func Add(a,b int) int {
        return a + b        
    }
    ```
    * 如何使用Go汇编实现一个函数？

2. 函数调用
    * 实参类型不一定要和形参一致: 可能存在隐式转换，形参是接口类型，调用函数时传入具体类型（只要能够隐式转换就没问题）
    * Go通过 __栈__ 传递函数的参数和返回值
        * 函数参数和返回值预先在栈上分配空间
        * 参数入栈从右到左，参数计算从左到右（何时计算）
        * 调用函数时参数都是传值（值复制，并且只复制值的直接部分）: 可以认为是标准变量声明
        * 函数参数不支持默认值（默认值是类型的零值）

3. 函数调用的退出阶段（返回开始到最终退出）
    * defer和go关键字差不多，go会将任务放入P队列，defer会将任务放入延迟队列
    * 两者放入队列时会计算出参数的值（参数优先计算，闭包的话会在执行时看使用的变量的值是啥）
    * defer所在的函数执行完成时，执行任务
```
type Foo struct {
	v int
}

func MakeFoo(n *int) Foo {
	print(*n)
	return Foo{}
}

func (Foo) Bar(n *int) {
	print(*n)
}

func main() {
	var x = 1
	var p = &x
	defer MakeFoo(p).Bar(p) // line 19 
	x = 2
	p = new(int) // line 21
	MakeFoo(p)
}
// Bar会被推入延迟队列: MakeFoo对于Bar来说只是个参数而已，其实就是数据相对于方法
// MakeFoo(p) -> 1
// MakeFoo(p) -> 这里的p是new(int)返回的地址，int默认零值为0
// Bar(p) -> 这里的p是&x，此时x为2
// 所以返回102
```

4. 匿名函数
    * 具名函数是匿名函数的一种特例
    * 在函数内部不能声明函数（匿名函数可以）
```
func squares() func() int {
    var x int
    return func() int {
        x++
        return x * x
    }
}

f := squares() // f不仅仅是函数（还记录了状态），这里的x和python类中的类属性有点像
fmt.Println(f())
fmt.Println(f())
```

5. 函数是 __一等公民__ 类型
    * 函数基本功能: 封装代码、分割功能、解耦逻辑
    * 作为普通的值: 在函数间传递、赋予变量、做类型判断和转换
        * 函数值不可被修改
    * 函数类型实现接口时，那么该接口类型包裹了一个函数，其实就可以把这个接口类型看成一个函数了
        * 接口不一定是数据类型，也可能是函数类型
    ```
    // 声明一个函数类型
    type Printer func(contents string)(n int, err error)
    
    func printToStd(contents string) (bytesNum int, err error){
        return fmt.Println(contents)
    }

    func main(){
        var p Printer
        p = printToStd // 把函数作为一个普通的值赋给一个变量
        p("something")
    }
    ```
    * 变长参数
    ```
    // 声明时
    func Sum(values ...int64) int64     // 变长参数总是一个切片类型 ...T == []T
                                        // 函数参数不能设置默认值[]
                                        // ...T 和 []T 还是有区别的: func (values []int) int
    func Println(a ...interface{})(n int, err error) //通过reflect来获取本来的值

    // 调用时
    int64_slice := []int64{1,2,3,4}
    Sum(int64_slice...)

    // fmt.Printf("%s\n", "hello")
    func errorf(linenum int, format string, args ...interface{}) {
        fmt.Fprintf(os.Stderr, "Line %d: ", linenum)
        fmt.Fprintf(os.Stderr, format, args...)
        fmt.Fprintln(os.Stderr)                    // 添加一个空格
    }
    ```

6. 什么是高阶函数？
    * 接受其他的函数作为参数
    * 把其他函数作为结果返回
    * 按功能来描述一下不同的函数
        * 普通函数: 添加处理逻辑
        * 工厂函数: 生成一个函数（返回一个函数）
        * 装饰器函数: 增强函数功能（参数是一个函数，返回一个函数）
        * handler函数: 额外添加自定义处理逻辑（参数是一个handler函数，返回值）
    ```
    // 对operate的约束
    type operate func(x,y int) int

    func caculate(x,y int, op operate) (int, error) {
        // 卫述语句: 用来检查关键的先决条件的合法性，并在检查未通过的情况下立即终止当前代码的执行
        if op == nil {
            return 0, errors.New("invalid operation")
        }
        return op(x,y), nil
    }
    ```

7. 如何实现闭包？
    * 闭包: 函数中存在对外来标识符的引用(也叫自由变量)
    * 这些自由变量作用域提升了，调用函数的时候，都能访问到这个公共变量
    * 有些变量在函数外部已经不使用了，函数 __包裹__ 了这些变量，延长了这些变量的声明周期
    ```
    // genCalculator -> 高阶函数
    // op -> 自由变量
    // calculateFunc -> 闭包函数的类型
    func genCalculator(op operate) calculateFunc {
        return func(x int, y int) (int, error) {
            // 这个函数中的op还不确定
            // 捕获自由变量，形成闭包函数
            if op == nil {
                return 0, errors.New("invalid operation")
            }
            return op(x, y), nil
        }
    }
    ```
    * 实现闭包的意义？
        * 动态生成部分逻辑
        * 模板方法

8. unsafe包中的函数调用是在编译时执行的（wc！）

9. 设计模式: Go函数参数不支持默认值，怎么处理的？
    ```
    type Server struct {
        Addr     string
        Port     int
        Protocol string
        Timeout  time.Duration
        MaxConns int
        TLS      *tls.Config
    }
    ```
    * 如何初始化不同的Server？
        1. 构造函数
            * 多写几个构造函数进行初始化
        ```
        func NewDefaultServer(addr string, port int)(*Server, error){
            return &Server{addr, port, "tcp", 30 * time.Second, 100, nil}, nil
        }
        ```
        2. 把必要参数和非必要参数分开
            * 那就不用写那么多构造函数了
            * 只需要写config配置
        ```
        type Config struct {
            Protocol string
            Timeout  time.Duration
            Maxconns int
            TLS      *tls.Config
        }

        type Server struct {
            Addr string
            Port int
            Conf *Config
        }

        func NewServer(addr string, port int, conf *Config) (*Server, error) {
            //...
        }

        //Using the default configuratrion
        srv1, _ := NewServer("localhost", 9000, nil) 

        conf := ServerConfig{Protocol:"tcp", Timeout: 60*time.Duration}
        srv2, _ := NewServer("locahost", 9000, &conf)
        ```
        3. Builder模式
            * 首先创建一个空的Server，然后调用方法添加属性，通过Build方法返回Server
            * 需要加什么属性，就加什么属性
        ```
        //使用一个builder类来做包装
        //不包装成builder就不行吗？
        type ServerBuilder struct {
            Server
        }

        func (sb *ServerBuilder) Create(addr string, port int) *ServerBuilder {
            sb.Server.Addr = addr
            sb.Server.Port = port
            //其它代码设置其它成员的默认值
            return sb
        }

        func (sb *ServerBuilder) WithProtocol(protocol string) *ServerBuilder {
            sb.Server.Protocol = protocol 
            return sb
        }

        func (sb *ServerBuilder) WithMaxConn( maxconn int) *ServerBuilder {
            sb.Server.MaxConns = maxconn
            return sb
        }

        func (sb *ServerBuilder) WithTimeOut( timeout time.Duration) *ServerBuilder {
            sb.Server.Timeout = timeout
            return sb
        }

        func (sb *ServerBuilder) WithTLS( tls *tls.Config) *ServerBuilder {
            sb.Server.TLS = tls
            return sb
        }

        // 返回内部Server
        func (sb *ServerBuilder) Build() (Server) {
            return  sb.Server
        }

        sb := ServerBuilder{}
        server, err := sb.Create("127.0.0.1", 8080).
            WithProtocol("udp").
            WithMaxConn(1024).
            WithTimeOut(30*time.Second).
            Build()
        ```
        4. Functional Options: 更易扩展
        ```
        type Option func(*Server) // 函数类型

        // 
        func Protocol(p string) Option { 
            return func(s *Server) { 
                s.Protocol = p 
            }
        }

        // Option和Middleware对比，我们需要的是一个数据类型，还是一个函数类型？
        // 数据类型: 我们需要初始化一个数据类型，然后通过Option来定义该数据类型
        // 函数类型: 我们需要增强一个函数类型的功能，然后通过Middleware来装饰该函数类型
        func NewServer(addr string, port int, options ...Option) (*Server, error) {
            // 1. 默认参数
            srv := Server{
                Addr:     addr,
                Port:     port,
                Protocol: "tcp",
                Timeout:  30 * time.Second,
                MaxConns: 1000,
                TLS:      nil,
            }
            // 2. 可选配置（通过传入函数来修改默认配置）
            for _, option := range options {
                option(&srv)
            }
            //...
            return &srv, nil
        }

        // 使用
        server := NewServer("192.168.1.1", "8080", Protocol("TCP"))
        ```

10. Go1.23新的语法糖: 遍历函数值
    * 下面3种函数类型可以通过for-range遍历
        * func(yield func() bool)
        * func(yield func(V) bool)
        * func(yield func(K, V) bool)
```
func Loop3(yield func() bool) {
	for range 3 {
		if (!yield()) {
			return
		}
	}
}

// 和python中的生成器差不多，主要是为了通过函数来生成值
for range Loop3 {

}

```
