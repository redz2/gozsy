# interface value
1. 接口类型的内部表示
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
2. 接口值
    * 动态类型 + 动态值
    * 通过fmt.Printf("%T\n", w)来获取动态值的类型（fmt内部使用reflect）

3. 接口类型的变量（接口类型的值能够包裹非接口类型的值）
    * nil error 不等于 nil？
    ```
    type MyInt interface {}
    var pa *int = nil     // 会将无类型的nil隐式转换为*int的nil（虽然pa == nil，但是pa此时已经有类型了）
    var myint MyInt = pa  // 何时真的是nil？未显式初始化，或者显式初始化为nil
    if myint == nil {     // myint接口是否包裹了其他值（显然这里包裹了pa，所以不为nil）
        ...
    }
    ```
    * 值包裹
        * 什么是接口: 一个用来包裹非接口类型值的 __盒子__
        * __接口赋值__ 的一些细节（如何包裹其他类型的值？）
            * 官方说法: 如果类型T实现了接口I，T可以 __隐式转换__ 成I
                * 这就是为什么函数参数使用接口，可以传入具体类型
                * 函数参数使用接口，其实就实现了多态
            * 如果T是非接口值（动态类型 + 动态值），T的复制会被包裹在结果I值中
                * 动态值的直接部分无法修改？？？
            * 如果T也是接口值，T包裹的值会复制一份到I值中
        ```
        // 一个*Book值被包裹在了一个Aboutable值中。
        var a Aboutable = &Book{"Go语言101"}

        // i是一个空接口值。类型*Book实现了任何空接口类型。
        var i interface{} = &Book{"Rust 101"}

        // Aboutable实现了空接口类型interface{}。
        i = a
        ```
        * fmt打印一个接口类型的变量
            * 输出动态值
        * Go构建一个表来记录各个类型的信息（类型种类，以及此类型的方法、字段、尺寸）
        * 接口包含什么？
            * 动态类型的信息: 动态类型 + 动态值 -> 实现反射
            * 一个方法列表（切片） -> 实现多态