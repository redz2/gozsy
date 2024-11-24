# 程序基本元素
* 操作
    * 函数调用
    * 运算符操作
    * 条件语句
    * 循环语句
* 数据
    * 类型（类型可以看作是值的模板，值可以看作类型的实例 -> 确定内存布局）
        * 内置类型
            * bool
            * string
            * int
                * type byte = uint8（类型种类Kind相同）
                * type rune = int32（任一类型的所有值的尺寸都是相同的，一般用字节来表示，rune32尺寸为4）
            * float
        * 自定义类型
    * 值
        * 字面量
            * 基本类型的字面量都是无名常量或者字面常量
                * 字面量都属于不确定类型
                    * 但是有默认类型，比如我写32（32本身并没有任何类型，只不过默认类型是int）
                * 类型不确定常量的显式转换T(v)
                ```
                int(1.23) // 非法转换
                int32(12)
                ```
                * 类型推断
                ```
                a := 97  // int
                b := 'a' // int32
                c := "a" // string
                ```
            * 字面量只是在代码中的文字体现（让编译器认为这么写就是某一种类型）
            * bool -> true/false
            * string
                * 直白 `abc\n`
                * 解释 "abc\n"
            * int
                * 十六进制 0x17
                * 十进制 17
                * 八进制 017 0o17 0O17
                * 二进制 0b1001 0B1001
            * rune
                * 可以用值表示
                * 字符 'a'
                * 变种表示: 
                    * 八进制 '\141'
                    * 十六进制 '\x61'
                    * '\u0061'
                    * 其他特殊字符
                        * '\t'
                        * '\a'
            * 分段提高可读性
                * 0x1_000_000
        * 变量
            * 变量是运行时存储在内存中可以修改的具名的值
            * 所有变量都是类型确定值
            * 变量声明（并初始化）
                * 标准变量声明
                ```
                var s string = "Go"
                var website = "baidu"       // 类型推断可以视为隐式类型转换
                var a int                   // 初始化为零值
                ```
                * 短变量声明
                ```
                c := 1
                ```
            * 赋值
                * 纯赋值语句
                    * 左边必须是一个可寻址的值、一个映射元素、一个空标识符
                    * 常量是不可寻址的
                * 短变量声明
        * 常量
            * 具名常量声明
            ```
            const Pi = 3.14159
            const X float32 = 3.14  // 类型确定具名常量

            const (
                X float32 = 3.14
                Y                   // 自动补全 Y float32 = 3.14
                Z
            )

            const (                 // iota初始为0，出现在第几个，就是几
                k = 3 // 在此处，iota == 0
                J = 3 // 在此处，iota == 0
                m float32 = iota + .5 // m float32 = 2 + .5
                n                     // n float32 = 3 + .5
            )
            ```
            * 常量是没有类型的
    * 问题
        * 每个类型的零值
        * type MyInt int MyInt和int类型种类都是int，但他两是两个不同的类型
* 关键字
* 标识符
    * Unicode字母或_开头（不能是数字）
    * Unicode大写字母开头 -> 导出标识符
* 表达式和语句
    * 一个表达式表示一个值，一条语句表示一个操作
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