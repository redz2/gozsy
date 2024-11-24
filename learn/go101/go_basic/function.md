# 函数
1. 函数声明
    * 可以把参数和结果声明看作是省略了var关键字的标准变量声明
    * Go不支持参数默认值
    ```
    func Add(a,b int) int {
        return a + b        
    }
    ```

2. 函数调用
    * 函数传参也属于赋值操作
    * 实参类型不一定要和形参一致（能够隐式转换）

3. 函数调用的退出阶段
    * 延迟函数

4. 匿名函数

5. 函数是一等公民
    * 函数基本功能: 封装代码、分割功能、解耦逻辑
    * 作为普通的值: 在函数间传递、赋予变量、做类型判断和转换
    ```
    // 声明一个函数类型
    type Printer func(contents string)(n int, err error)
    
    func printToStd(contents string) (bytesNum int, err error){
        return fmt.Println(contents)
    }

    func main(){
        var p Printer
        // 把函数作为一个普通的值赋给一个变量
        p = printToStd
        p("something")
    }
    ```
    * 变长参数
    ```
    func (values ...int64) int64      // 变长参数总是一个切片类型 ...T == []T
                                      // 调用函数时，使用方式上有区别  
    func Println(a ...interface{})(n int, err error) //通过reflect来获取本来的值
    ```

6. 什么是高阶函数？ --- 让函数在其他函数之间传递
    * 接受其他的函数作为参数
    * 把其他函数作为结果返回
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
    * 闭包: 函数中存在对外来标识符的引用(自由变量)
        * 体现从“不确定”变成“确定”的过程
    * 自由变量: 提升变量作用域
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

8. 所有函数调用的传参都是值复制
    * 只复制值的直接部分

9. 一个原则
    * 既不要把你程序的细节暴露给外界，也尽量不要让外界的变动影响到你的程序

10. unsafe包中的函数调用是在编译时执行的（wc！）

11. Go两种设计模式
    * 函数选项模式: 结构体创建（解决函数不支持默认值）
    ```
    type Server struct {
        addr string
        port string
    }

    type Option func(*Server)

    // 返回一个函数
    func WithAddr(addr string) Option {
        return func(s *Server){
            s.addr  = addr
        }
    }

    func NewServer(options ...Option) *Server {
        srv := &Server{
            addr: "localhost",
            port: "8080",
        }
        for option := range options {
            // 执行函数修改参数
            option(srv)
        }
    }

    // 如何使用
    NewServer(WithAddr("192.168.1.1"))
    ```
    * 装饰器模式（用于中间件）
    ```
    // 函数签名一定要一模一样
    // 如何增强函数功能？
    func Logger(handler Handler) Handler{
        return func(w http.ResponseWrite, r *http.Request){
            now := time.Now()
            handler(w, r)
            log.Printf("url: %s, time: %v", r.URL, time.Since(now))
        }
    }
    ```
