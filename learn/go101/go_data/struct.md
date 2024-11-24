# 结构体
* 结构体类型和结构体字面量表现形式
    * 无名结构体类型的字面形式
        * 为了代码可读性，最好少使用匿名结构体类型
    ```
    struct {
        name string // 成员变量，字段
        age  int
    }
    ```
    * 结构体类型的尺寸
        * 所有字段尺寸加上一些填充字节（字节填充、内存地址对齐），如何进行填充？
            * 64位机器一次性可以读取8字节
            * 逻辑上连续的8个字节其实不连续（并行读取8字节，提高效率）
                * 如果不是逻辑上连续的字节，会分多次读取（连续的字节可以一次读取）
                * 编译器为了提高性能，会做内存对齐（能一次读取所有数据就一次性读取所有数据）
                    * example: int64类型的值长8字节，要保证能一次读取
            * 各种数据类型的内存大小（为了可以一次读取）
                * int8   1byte
                * int16  2byte
                * int64  8byte
                * string 16byte
                * slice  24byte
                ```
                type T struct {
                    a int8     
                    b int64
                    c int32
                    d int16
                }
                // 起始地址为0，其他的都是相对地址
                // 整体必须是字节对齐的倍数
                x0000000 xxxxxxxx xxxxxx00
                ```
        * 零字段结构体尺寸为零

    * StructTag: 结构体标签
        * 通过reflect.StructTag解析tag内容
        ```
        type Example struct {
            Foo string `json:"foo,Omitempty" xml:"foo"` // 形式为key:"value,option"
        }
        tag := reflect.StructTag(`json:"foo,omitempty"`)  // 使用reflect.StructTag解析tag内容
        value := tag.Get("json")
        fmt.Printf("value: %q\n", value)
        ```
        * 通过reflect获取结构体的tag
        ```
        // 可以获取struct对应字段的json的tag
        type Vehicle struct {
            ID       int    `json:"id"`
            CityName string `json:"city_name"`
        }

        func main() {
            reflectType := reflect.ValueOf(Vehicle{}).Type()
            fmt.Printf("fileds number: %v\n", reflectType.NumField())
            for i := 0; i < reflectType.NumField(); i++ {
                fmt.Printf("%v", reflectType.Field(i).Name)
                // reflect.StructField包含一个字段reflect.StructTag
                fmt.Printf(" tag:%v\n", reflectType.Field(i).Tag.Get("json"))
            }
        }
        ```
    * 结构体字段不能包含自身
    * 组合字面量形式: T{...}
        * 什么是组合？
    * 结构体的赋值
        * 等价于每个字段逐个赋值

* 结构体的可寻址性
    ```
    var book = Book{} // 变量可寻址

    Book{}.title      // 字面量不可寻址（思考下变量和字面量的区别，编译器会推断字面量的类型，字面量没有类型）
    p := &Book{100}   // 语法糖: 组合字面量可寻址 <=> tmp := Book{100}; p := &tmp

    (*bookN).page     // bookN.page，通过结构体指针进行字段选择时，会隐式解引用
    ```

* 方法 vs 函数
    * 函数: 可以有名字，可以没有名字，可以被当成值传递
    * 方法: 需要有名字，不能被当做值来看待，必须隶属于一个类型
        * 方法可以隶属于任何自定义类型，但不能是接口类型
        * 一个类型关联的所有方法，组成了该类型的方法集合
        * 方法集合中的方法不能重名，与字段名称也不能重复 ***
        * 类型的一个字段看作是一个属性，一个方法看作是一个能力，将属性和能力封装在一起，就是面向对象编程的一个主要原则
        * 多使用组合，而不是使用继承
    ```
    type Person struct {
        name string
        age  int
    }

    func (p Person) String() string{
        // 可以使用当前值的任何一个字段，任何一个方法(包括String方法本身)
        return fmt.Sprintf("%s %d", p.name, p.age)
    }

    p := Person{name: "zhouyi"}
    fmt.Printf("%s\n", p)
    ```

* 结构体的嵌入字段
    * 嵌入字段(匿名字段)既是类型，也是名称
    * 多使用组合（不要用继承），通过组合的方式来丰富该类型的属性和功能
    ```
    type Student struct {
        id       int
        Person
        Singer
        Dancer
    }

    // 嵌入字段的“字段”和“方法”都会被“嫁接”到结构体中
    // 嵌入字段的“字段”和“方法”，都会被结构体中同名的“字段”和“方法”屏蔽
    // 即使屏蔽了，我们也可以通过链式的选择表达式，选择嵌入字段的“字段”和“方法”
    func (s Student) String () string{
        return fmt.Sprintf("id: %d, Person: %s", s.id, s.Person) // “逐层包装”，这里会调用Person的String方法
    }
    ```
    * 结构体的多层嵌入
        * 按照每一层去找字段或方法，嵌入层级越深的字段或方法越可能被“屏蔽”
        * 如果同一层级的多个嵌入字段拥有同名字段的字段或方法，编译器也不知道该选哪个，会报编译错误

* golang虽然支持面向对象编程，但没有继承，只有组合
    * 通过组合来扩充功能
    * 通过名称屏蔽来重写功能

* 值方法和指针方法
    * 值方法的接收者是该值的副本（对值修改不会改变原值，除非本来就是引用类型）
    * 指针方法接收的是该值的指针值的副本（修改一定会体现在原值上）
    * 一个自定义类型是否实现了某个接口？
        * 值类型和指针类型的方法集合不同
            * 一个自定义类型的方法集中包含所有值方法
            * 一个自定义类型的指针类型方法集包含所有值方法和指针方法
        * 第二步，看包含的方法是否能覆盖接口中的方法集
    * 一个自定义类型能否调用指针方法？
        * 当值的方法集中没有方法时，会去指针的方法集中找（指针本身包含值方法和指针方法）
        * 可以，go编译器会自动做转换 cat.SetName("xxx") -> (&cat).SetName("yyy")
            * 算是语法糖，或者编译器认为你写错了，帮你改过来了
        * 虽然能调用指针方法，该类型并未实现包含SetName方法的接口（有一点点矛盾）
    ```
    func (cat Cat) GetName()string {
        return cat.name
    }

    func (cat *Cat) SetName(name string) {
        cat.name = name
    }
    ```