# 类型系统
## 问题
1. 什么叫类型？
2. 什么叫值？
    * 一个类型的一个实例称为此类型的一个值

## 类型（Type）
1. 基本类型(预声明类型)
    * 字符串: string
    * 布尔: bool
    * 数值类型: int,float,complex
        * byte是uint8的类型别名，rune是int32的类型别名
        ```
        // type后有没有=差别巨大
        type byte = uint8    // 类型别名: byte和uint8完全等价，byte只是一个名称，更明确地表示这代表一个字符
        type MyInt int       // 自定义类型: MyInt和int完全是两个类型，只是底层数据存储一样
        ```
2. 组合类型
    * 指针类型: *T
    * 结构体类型:
    ```
    struct {
        name string
        age  int
    }
    ```
    * 函数类型: func(int,int)int
    * 容器类型
        * 数组: [5]T
        * 切片: []T
        * map: map[Tkey]T
    * 通道类型
    ```
    chan T
    chan<- T
    <-chan T
    ```
    * 接口类型 - 反射和多态
    ```
    type interface {
        Method0(string)int
        Method1()(int, error)
    }
    ```
3. 类型种类（what is Kind？）
    * Name: 类型名称为User
    * Kind: 类型种类为struct(go一共26种类型种类)
    ```
    type User struct {
        name string
        age  int
    }
    ```
4. 类型定义（what is type？类型定义声明）
    * 使用type关键字定义一个“新”类型
    ```
    type NewTypeName SourceType // 这是两个类型，底层类型一致，可以显式类型转换
    ```
    * 具名类型和无名类型
        * 具名类型: type MyInt int
        * 无名类型: 一定是组合类型（反之则未必）
        ```
        x := struct{
            name string
        }{
            "shaw",
        }
        ```
        * 这说明一个类型一定有种类，但不一定会有名字
    * 类型别名
        * byte和uint8是相同的类型
    * 底层类型（比较重要的概念）
        * 一个内置类型的底层类型是自己: int
        * 一个无名类型的底层类型是自己: 
        * 新声明的类型和源类型共享底层类型: type MyInt int
        * 如何溯源一个类型的底层类型？
            * 遇到内置类型或无名类型结束
            ```
            type MyInt int  // 底层类型为int
            type MyInts []MyInt // 底层类型为[]MyInt（slice）
            ```
        * 底层类型有啥用？
            * 类型转换、赋值、比较

## 值（Value）
1. 值在代码中以什么形式呈现？
    * 字面量
        * 指针类型、通道类型、接口类型没有字面量表示形式
        * 值分为类型确定的和类型不确定的
    * 具名常量
    * 变量
    * 表达式
2. 值部
    * Go类型分为两大类别（按内存分布分）
        * 每个值在内存中只分布在一个内存块上的类型（直接值部）
            * bool
            * int
            * ptr
            * struct
            * array
        * 每个值在内存中会分布在多个内存块上的类型（直接值部 -> 底层间接值部）
            * slice
            * map
            * channel
            * func
            * interface
            * string(值类型)
            ```
            type _map *hashtableImpl
            type _channel *channelImpl
            type _function *functionImpl

            type _slice struct {
                elements unsafe.Pointer // 引用着底层的元素
                len      int            // 当前的元素个数
                cap      int            // 切片的容量
            }

            type _string struct {
                elements *byte // 引用着底层的byte元素
                len      int   // 字符串的长度
            }

            type _interface struct {
                dynamicType  *_type         // 引用着接口值的动态类型
                dynamicValue unsafe.Pointer // 引用着接口值的动态值
            }
            ```
        * 相关问题
            * 什么是引用类型呢？这个概念没有必要存在（使用“指针持有者类型”来替代引用类型）
                * 包含指针的一定是引用类型吗？一般来说是的，但string就是值类型（编译器不让修改）
            * 赋值时，底层间接值部不会被复制（只复制直接值部，共享底层间接值部）
                * 这就是为啥说引用类型的赋值是浅拷贝的原因，go中没有浅拷贝、深拷贝这些概念，如果一定要有，都是浅拷贝
                * 字符串和接口可能不太一样，接口的动态值是只读的，字符串也是只读的，即使共享底层数据，因为无法修改，就不存在传引用的修改值的问题
    * 值尺寸
        * unsafe.Sizeof: 一般说的是直接部分的字节数，就是类型的尺寸
    * 一个值的具体类型和具体值
        * 非接口类型就是值的类型和值本身
        * 接口类型具体类型和具体值指的是动态类型和动态值

## nil
1. nil是一个预声明的标识符（可以表示多种类型的零值，他可以表示很多内存布局不同的值，而不仅仅是一个值）
    * 可以表示以下种类（Kind）的类型的零值（虽然零值都用nil表示，但他们的零值并不相同）
        * 指针
        * map
        * slice
        * func
        * channel
        * interface
    * nil没有默认类型（不是没有类型）
        * true/fasle: bool
        * iota: int
        * 必须为编译器提供信息能判断出nil的类型
    * 不同种类的类型的nil值的内存大小不同
        * 同一类型的nil和非nil的值内存大小是一致的（nil值只是一种特殊表示）
            * 通过unsafe.Sizeof()可以证明nil和非nil的内存结构一致
        * 不同种类的nil无法比较(因为类型不同)
            * (interface{})(nil) == (*int)(nil)  // 可以比较，但是不相等（接口中包裹了什么？还是什么都没包裹）
                                                 // 一个nil接口什么也没包裹，另外一个会转换成包裹了nil非接口值的接口值
        * map、func、slice不支持比较，只能和nil进行比较
2. nil: 提供了一个关键词用来表示一些类型的零值