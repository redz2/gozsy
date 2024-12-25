# 程序源代码基本元素
* 编程可以看作是以各种方式控制和组合计算机运行中的各种操作
    * 最基本的操作就是底层计算机指令 -> 高级编程语言进行抽象和封装
    * 那在高级语言中一个操作是什么？
        * 函数调用: Sum(1, 2)
        * 使用运算符运算: 1 + 2
        * 条件语句
        * 循环语句

* 操作啥？数据？那什么是数据？
    * 类型 -> 可以看作是值的模板
        * type 类型名称 底层类型
        * 类型转换不会改变值本身，但是语义会发生变化（只改变类型）
    * 值 -> 可以看作是类型的实例
        * 字面量 -> 编码时如何书写？类型不确定，但有默认类型
            * 一个值如何表示和一个值实际是什么有啥区别？
            * 比如我写1，一定是int吗？也可能是int64，只不过默认是int而已
        * 变量 -> 类型确定，值不确定
            * var 变量名字 变量类型 = 表达式
            * new(T): 创建一个T类型的匿名变量，初始化为T类型的零值，返回变量地址*T
            * 变量的生命周期？
            * 变量是否从函数中逃逸: 是否还需要这个变量？
        * 具名常量

* 在一个源代码文件中我们能看到哪些东西？
    * Go关键字: func var ...
    * Go标识符
        * 代码要素: 具名的函数、具名的值、定义的类型、包名
            * 代码要素名就是标识符
        * Unicode字母或_开头（开头不能是数字）
        * Unicode大写字母开头（包外可见）

* 表达式和语句
    * 一个表达式表示一个值，一条语句表示一个操作
    ```
    // 理解表达式的求值顺序
    // 1. 包级别变量声明语句
    var (
        a = c + b  // 如果某个变量依赖其他变量，那么初始化放在后面(包级别变量才行，函数中必须按顺序)
        b = f()
        c = f()
        d = 3
    )

    func f() int {
        d++
        return d
    }

    func main(){
        fmt.Println(a,b,c,d)
    }

    // 2. 普通求值顺序: 从左到右(一般都是这种情况，而不是包级变量)

    // 3. 赋值语句的求值
    n0, n1 := 1, 2
    n0, n1 = n0 + n1, n0  // 对于表达式从左到右进行求值，然后从左到右进行赋值

    // 4. switch/select语句中的表达式求值
    switch Expr(2) {
        case Expr(1),Expr(2),Expr(3):   // 惰性求值，从左到右依次求值，如果匹配成功，后面的表达式就不会进行求值了
            fmt.Println("enter into case1")
            fallthrough
        case Expr(4):
            fmt.Println("enter into case2")
    }
    ```
    * 基本表达式(单值表达式)
        * 字面量、变量、具名常量: 123
        * 预算符操作: 1 + 2
        * 一个函数返回一个值，则函数的调用属于表达式: f()
        * 通道的接收: <-c
    * 简单语句
        * 短变量声明
        * 纯赋值语句
        * 有返回结果的函数或方法调用
        * 通道的发送操作: c<-
        * 空语句
        * 自增或自减（不是表达式）
    * 一些非简单语句
        * 标准变量声明: var a int = 1
        * 具名常量声明: const b int =2
        * 类型声明: type MyInt int
        * 包引入语句: import "fmt"
        * 显式代码块: {}
        * 函数声明: func Add(a,b int)int {}
        * 流程控制: for、if
        * 函数返回语句: return
        * 协程创建: go Hello()
    * 非简单语句中有4种声明语句（var、const、type、func）
* 基本流程控制语法
    * 基本流程控制代码块
        * 条件: if-else
        ```
        if err := SomeFunc(), err != nil {
            ...
        } else {
            ...
        }
        ```
        * 循环: for
        ```
        for i := 0; i < 10; i++ {
            fmt.Println(i)
        }

        var i = 0
        for ; i < 10; {  // 如果中间的条件判断被省略，编译器视为true
            fmt.Println(i)
            i++
        }
        ```
        * 多条件: switch-case
        ```
        switch time.Now().Weekday() {         // 明天是星期几: time.Now().Add(time.Hour*24).Weekday()
        case time.Saturday, time.Sunday:
            fmt.Println("It's the weekend")
        default:
            fmt.Println("It's a weekday")
        }
        ```
    * 特殊
        * 用来遍历整数、各种容器、通道和某些函数的for-range循环代码块
        ```
        for i := range anInteger {  // 等价于for i := 0; i < anInteger; i++ {...}
            ...
        }
        ```
        * 接口相关的type-switch多条件分支代码块
        ```
        whatAmI := func(i interface{}) {
            switch t := i.(type) {
            case bool:
                fmt.Println("I'm a bool")
            case int:
                fmt.Println("I'm an int")
            default:
                fmt.Printf("Don't know type %T\n", t)
            }
        }

        whatAmI(true)
        ```
        * 通道相关的select-case多分支代码块
        * goto语句
        ```
        Next: // 跳转标签声明
            fmt.Println(i)
            i++
            if i < 5 {
                goto Next // 跳转
            }
        ```