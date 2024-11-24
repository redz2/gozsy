# 反射
* 反射: 在运行期探知对象的类型信息和内存结构
    * 剖析一个变量（从类型到值）
    * 两个反射入口函数
    ```
    func TypeOf(i interface{})Type      // 返回一个变量的类型
    func ValueOf(i interface{})Value    // 返回一个变量的值
    ```
* 类型
    * Type获取数据类型信息
    ```
    // Name是类型名称，Kind是类型的底层类型
    type X int
    var a X = 1
    at := reflect.TypeOf(a)
    fmt.Println(at.Name(), at.Kind()) // X int

    b := 100
    bt := reflect.TypeOf(b)
    fmt.Println(bt.Name(), bt.Kind()) // int int

    type User struct {
        Name string
        age  int
    }

    // 通过反射获取变量类型
    user := User{"xxx", 23}
    userType = reflect.TypeOf(user)
    userType.Name()     // 类型名称: User
    userType.Kind()     // 类型种类: struct
    userType.NumField() // 类型字段的个数
    userType.Field(i)   // 类型字段
    fieldType.Name()    // 类型每个字段名称
    fieldType.Tag()     // 类型每个字段标签

    // 如果变量是指针怎么办？如何获取指针指向对象的类型？
    user2 := &User{"yyy", 25}
    ptr := reflect.TypeOf(user2) // 类型就是ptr
    elem := reflect.TypeOf(user2).Elem() // 可以得到指针指向的元素的类型和名称
    ```
* 值
    * Value专注于对象实例的数据读写
    ```
    type User struct {
        Name string
        age  int
    }
    user := User{"xxx", 23}

    elem = reflect.ValueOf(user)
    elem.NumField()
    field := elem.Filed(i) // 字段的值


    // 接口变量会复制对象，且是unaddressable的（反正是找不到原来的内存空间的）

    // Value数据类型包含一个接口类型，接口类型包含我们传入的变量
    // 如果变量是指针，那我们有该变量指针的副本，通过指针就可以知道变量的类型和值
    // 如果变量不是指针，那我们只有该变量的副本，当然我们也可以知道变量的类型和值，但是我们不能修改该变量的值

    // 我认为Elem()的作用，标记reflect.Value包含的是指针
    // 如果调用SetInt()方法，需要额外调用*ptr
    // 如果不标记，鬼知道你是不是指针，肯定不会做额外处理
    // 对于类型也是一样的，指针需要额外处理
    a := 10
    va, vp := reflect.ValueOf(a), reflect.ValueOf(&a).Elem() // reflect.Value
    fmt.Println(va.CanAddr(), va.CanSet())
    fmt.Println(vp.CanAddr(), vp.CanSet())

    vp.SetInt()
    ```
* 动态调用方法
    * 
* 性能
    * 反射非常影响性能


