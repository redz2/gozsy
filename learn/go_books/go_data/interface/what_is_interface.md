# what is interface
1. 接口类型 vs 非接口类型（具体类型）
    * 当我拿到一个具体类型，我就知道这是什么东西，能用来做什么？
    * 当我拿到一个接口类型，我不知道这是什么，只知道通过它的方法来做什么

2. 接口能用来做什么？
    * 使用接口(interface)来实现装盒(value boxing)和反射(reflection)
    * 当接口作为函数参数使用（LSP: 里氏替换）
        * fmt.Fprintf: 输出到io.Writer(可以写入bytes)
        * fmt.Sprintf: 返回一个字符串
        * fmt.Printf: 输出到stdout
    * 接口类型: fmt.Stringer
    ```
    type Stringer interface {
        String() string
    }
    ```
